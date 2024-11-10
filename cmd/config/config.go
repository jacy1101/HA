package config

import (
	"HA/cmd"
	"HA/utils/Cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置云服务商的访问密钥 (Configure the cloud service provider's access key)",
	Long:  "配置云服务商的访问密钥 (Configure the cloud service provider's access key)",
	Run: func(cmd *cobra.Command, args []string) {
		Cmd.ConfigureAccessKey()
	},
}
