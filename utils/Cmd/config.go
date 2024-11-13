package Cmd

import (
	"HA/utils"
	"HA/utils/Cloud"
	"HA/utils/Error"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"strings"
)

const (
	alibaba = "alibaba"
	tencent = "tencent"
	//aws     = "aws"
	//huawei  = "huawei"
)

func ConfigureAccessKey() {
	cloudConfigList, cloudProviderList, cloudProvider := selectProvider()
	for i, j := range cloudProviderList {
		if j == cloudProvider {
			var credList []Cloud.Config
			switch cloudConfigList[i] {
			case alibaba:
				credList = append(credList, findAlibabaConfig()...)
			case tencent:
				credList = append(credList, findTencentConfig()...)
				//case aws:
				//	credList = append(credList, findAWSConfig()...)
				//case huawei:
				//	credList = append(credList, findHuaweiConfig()...)
			}
			if len(credList) != 0 {
				var (
					isTrue     bool
					selectedAK string
				)
				prompt := &survey.Confirm{
					Message: "在当前系统中发现访问密钥，选择是否使用该密钥？",
					Default: false,
				}
				err := survey.AskOne(prompt, &isTrue)
				Error.HandleFatal(err)
				if isTrue {
					var accessKeyList []string
					for i, v := range credList {
						i = i + 1
						accessKeyList = append(accessKeyList, strconv.Itoa(i)+"\t"+v.Provider+"\t"+v.Alias+"\t"+v.AccessKeyId)
					}
					accessKeyList = append(accessKeyList, "退出")
					sort.Strings(accessKeyList)
					prompt := &survey.Select{
						Message: "选择您使用的密钥：",
						Options: accessKeyList,
					}
					err := survey.AskOne(prompt, &selectedAK)
					Error.HandleFatal(err)

					if selectedAK == "退出" {
						log.Infoln("已退出")
						log.Exit(0)
						log.Debugln("正在退出……")
					} else {
						for _, v := range credList {
							if v.AccessKeyId == strings.Split(selectedAK, "\t")[3] {
								var config Cloud.Config
								config.Provider = tencent
								config.AccessKeyId = v.AccessKeyId
								config.AccessKeySecret = v.AccessKeySecret
								inputAccessKey(config, cloudConfigList[i])
							}
						}
					}
				} else {
					log.Infoln("已取消自动导入，请输入您要添加的访问密钥。")
					config := Cloud.Config{}
					inputAccessKey(config, cloudConfigList[i])
				}
			} else {
				config := Cloud.Config{}
				inputAccessKey(config, cloudConfigList[i])
			}
		}
	}
}

func selectProvider() ([]string, []string, string) {
	var cloudProvider string
	cloudConfigList, cloudProviderList := ReturnCloudProviderList()
	prompt := &survey.Select{
		Message: "选择您要设置的云服务商：",
		Options: cloudProviderList,
	}
	err := survey.AskOne(prompt, &cloudProvider)
	Error.HandleError(err)
	return cloudConfigList, cloudProviderList, cloudProvider
}

func ReturnCloudProviderList() ([]string, []string) {
	var (
		cloudConfigList   []string
		cloudProviderList []string
		CloudProviderMap  = utils.CloudProviderMap
	)
	for k, v := range CloudProviderMap {
		cloudConfigList = append(cloudConfigList, k)
		cloudProviderList = append(cloudProviderList, v)
	}
	return cloudConfigList, cloudProviderList
}

func inputAccessKey(config Cloud.Config, provider string) {
	OldAccessKeyId := ""
	OldAccessKeySecret := ""
	OldRegion := ""
	AccessKeyId := config.AccessKeyId
	AccessKeySecret := config.AccessKeySecret
	Region := config.Region
	if AccessKeyId != "" {
		OldAccessKeyId = fmt.Sprintf(" [%s] ", utils.MaskAK(AccessKeyId))
	}
	if AccessKeySecret != "" {
		OldAccessKeySecret = fmt.Sprintf(" [%s] ", utils.MaskAK(AccessKeySecret))
	}
	if Region != "" {
		OldRegion = fmt.Sprintf(" [%s] ", Region)
	}
	var qs = []*survey.Question{
		{
			Name:   "AccessKeyId",
			Prompt: &survey.Input{Message: "输入访问密钥 ID" + OldAccessKeyId + ":"},
			Validate: func(val interface{}) error {
				str := val.(string)
				if len(strings.TrimSpace(str)) == 0 && config.AccessKeyId == "" {
					log.Warnln("输入的 AccessKey 有误。")
				} else if len(strings.TrimSpace(str)) == 0 && config.AccessKeyId != "" {
					str = config.AccessKeyId
				}
				return nil
			},
		},
		{
			Name:   "AccessKeySecret",
			Prompt: &survey.Password{Message: "输入访问密钥密钥" + OldAccessKeySecret + ":"},
			Validate: func(val interface{}) error {
				str := val.(string)
				if len(strings.TrimSpace(str)) == 0 && config.AccessKeySecret == "" {
					log.Warnln("输入的 AccessKeySecret 有误。")
				} else if len(strings.TrimSpace(str)) == 0 && config.AccessKeySecret != "" {
					str = config.AccessKeySecret
				}
				return nil
			},
		},
		{
			Name:   "Region",
			Prompt: &survey.Input{Message: "输入凭证的地区（可选，默认为 ap-shanghai）" + OldRegion + ":"},
		},
	}
	cred := Cloud.Config{}
	err := survey.Ask(qs, &cred)
	cred.AccessKeyId = strings.TrimSpace(cred.AccessKeyId)
	cred.AccessKeySecret = strings.TrimSpace(cred.AccessKeySecret)
	cred.Region = strings.TrimSpace(cred.Region)
	cred.Provider = provider
	if cred.AccessKeyId == "" {
		cred.AccessKeyId = AccessKeyId
	}
	if cred.AccessKeySecret == "" {
		cred.AccessKeySecret = AccessKeySecret
	}
	if cred.Region == "" {
		cred.Region = "ap-shanghai"
	}
	Error.HandleError(err)
	log.Info(cred)
}
