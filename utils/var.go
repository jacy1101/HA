package utils

var CloudProviderMap = map[string]string{
	//"aws":     "AWS (Amazon Web Services)",
	"alibaba": "阿里云 (Alibaba Cloud)",
	"tencent": "腾讯云 (Tencent Cloud)",
	//"huawei":  "华为云 (Huawei Cloud)",
}

const (
	AppDirName = ".config/ha"
)

var TencentRegion = []string{"ap-guangzhou", "ap-shanghai", "ap-nanjing", "ap-beijing", "ap-chengdu", "ap-chongqing", "ap-hongkong", "ap-singapore", "ap-jakarta", "ap-seoul", "ap-tokyo", "ap-bangkok", "sa-saopaulo", "na-siliconvalley", "na-ashburn", "eu-frankfurt"}

var AlibabaRegion = []string{"ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-5", "ap-southeast-6", "ap-southeast-7", "ap-northeast-1", "ap-northeast-2", "us-west-1", "us-east-1", "eu-central-1", "eu-west-1", "me-east-1", "me-central-1", "cn-hangzhou-finance", "cn-shanghai-finance-1", "cn-shenzhen-finance-1", "cn-beijing-finance-1", "cn-north-2-gov-1"}
