package Cmd

import (
	"HA/utils"
	"HA/utils/Cloud"
	"github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
)

// 获取腾讯云配置文件
func findTencentConfig() []Cloud.Config {
	var credList []Cloud.Config
	// 1.credential file
	tencentConfigPath := filepath.Join(utils.GetUserDir(), "/.tccli")
	tencentConfigFiles, _ := os.ReadDir(tencentConfigPath)
	for _, f := range tencentConfigFiles {
		tencentConfigName := f.Name()
		if path.Ext(tencentConfigName) == ".credential" {
			tencentConfigFile := filepath.Join(tencentConfigPath, tencentConfigName)
			isTrue, content := utils.ReadFile(tencentConfigFile)
			if isTrue {
				contentJson, _ := simplejson.NewJson([]byte(content))
				cred := Cloud.Config{}
				cred.AccessKeyId = contentJson.Get("secretId").MustString()
				cred.AccessKeySecret = contentJson.Get("secretKey").MustString()
				cred.Provider = tencent
				if cred.AccessKeyId != "" {
					credList = append(credList, cred)
				}
			}
		}
	}

	// 2.environment variables
	cred := Cloud.Config{}
	cred.Provider = tencent
	accessKey := os.Getenv("TENCENTCLOUD_ACCESS_KEY_ID")
	if utils.IsValidSecretAccessKey(accessKey) {
		cred.AccessKeyId = accessKey
	} else {
		cred.AccessKeyId = ""
	}
	accessKeySecret := os.Getenv("TENCENTCLOUD_ACCESS_KEY_SECRET")
	if utils.IsValidSecretAccessKey(accessKeySecret) {
		cred.AccessKeySecret = accessKeySecret
	} else {
		cred.AccessKeySecret = ""
	}
	region := os.Getenv("TENCENTCLOUD_REGION")
	if utils.IsValidRegion(utils.TencentRegion, region) {
		cred.Region = region
	} else {
		cred.Region = ""
	}

	credList = append(credList, cred)

	log.Infoln(credList)
	return credList
}

// 获取阿里云配置文件
func findAlibabaConfig() []Cloud.Config {
	var credList []Cloud.Config
	// 1.credential file
	alibabaConfigFile := filepath.Join(utils.GetUserDir(), "/.aliyun/config.json")
	isTrue, content := utils.ReadFile(alibabaConfigFile)
	if isTrue {
		contentJson, _ := simplejson.NewJson([]byte(content))
		contentJsonArray, _ := contentJson.Get("profiles").Array()
		for _, v := range contentJsonArray {
			cred := Cloud.Config{}
			contentResult, _ := v.(map[string]interface{})
			cred.AccessKeyId = contentResult["access_key_id"].(string)
			cred.AccessKeySecret = contentResult["access_key_secret"].(string)
			cred.Provider = alibaba
			if cred.AccessKeyId != "" {
				credList = append(credList, cred)
			}
		}
	}
	// 2.environment variables
	cred := Cloud.Config{}
	cred.Provider = alibaba
	accessKey := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	if utils.IsValidSecretAccessKey(accessKey) {
		cred.AccessKeyId = accessKey
	} else {
		cred.AccessKeyId = ""
	}
	accessKeySecret := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	if utils.IsValidSecretAccessKey(accessKeySecret) {
		cred.AccessKeySecret = accessKeySecret
	} else {
		cred.AccessKeySecret = ""
	}
	region := os.Getenv("ALIBABA_CLOUD_REGION")
	if utils.IsValidRegion(utils.AlibabaRegion, region) {
		cred.Region = region
	} else {
		cred.Region = ""
	}
	if cred.AccessKeyId != "" {
		credList = append(credList, cred)
	}
	return credList
}
