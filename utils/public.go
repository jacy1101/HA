package utils

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetUserDir() string {
	home, _ := os.UserHomeDir()
	return home
}

func GetConfigFilePath() string {
	home, _ := GetHAHomeDir()
	CreateFolder(home)
	configFilePath := filepath.Join(home, "cache.db")
	return configFilePath
}

func GetHAHomeDir() (string, error) {
	return filepath.Join(GetUserDir(), AppDirName), nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CreateFolder(folder string) {
	if !FileExists(folder) {
		log.Tracef("创建 %s 目录 (Create %s directory): ", folder, folder)
		_ = os.MkdirAll(folder, 0700)
	}
}

func ReadFile(filePath string) (bool, string) {
	if FileExists(filePath) {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		return true, string(content)
	} else {
		return false, ""
	}
}

func MaskAK(ak string) string {
	if len(ak) > 7 {
		prefix := ak[:2]
		suffix := ak[len(ak)-6:]
		return prefix + strings.Repeat("*", 18) + suffix
	} else {
		return ak
	}
}

// 检查 Secret Access Key 格式
func IsValidSecretAccessKey(key string) bool {
	// 正则表达式：由 40 个字符组成，允许字母、数字
	regex := `^[A-Za-z0-9]{7,}`
	matched, err := regexp.MatchString(regex, key)
	if err != nil {
		return false
	}
	return matched
}

// 检查 Region 是否合理
func IsValidRegion(regions []string, region string) bool {
	for _, r := range regions {
		if r == region {
			return true
		}
	}
	return false
}
