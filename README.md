## HA 支持 Tencent Cloud && Alibaba Cloud

- 读取 1password 里的 AccessKey ，安装 [CLI](https://developer.1password.com/docs/cli/get-started)

```shell
export TENCENT_CLOUD_ACCESS_KEY_ID="op://app-prod/Tencent/access key id"
export TENCENT_CLOUD_ACCESS_KEY_SECRET="op://app-prod/Tencent/secret access key"
export TENCENT_CLOUD_REGION="op://app-prod/Tencent/xxxxxxxxxxxxxx"

op run --no-masking -- go run main/main.go config
```