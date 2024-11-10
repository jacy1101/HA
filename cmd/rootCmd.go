package cmd

import (
	"HA/utils"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"os"
)

var logLevel string

var RootCmd = &cobra.Command{
	Use:   "ha",
	Short: "Honey AccessKey 是一个快速创建蜜标的工具。",
	Long: `
██╗  ██╗ ██████╗ ███╗   ██╗███████╗██╗   ██╗     █████╗ ██╗  ██╗
██║  ██║██╔═══██╗████╗  ██║██╔════╝╚██╗ ██╔╝    ██╔══██╗██║ ██╔╝
███████║██║   ██║██╔██╗ ██║█████╗   ╚████╔╝     ███████║█████╔╝ 
██╔══██║██║   ██║██║╚██╗██║██╔══╝    ╚██╔╝      ██╔══██║██╔═██╗ 
██║  ██║╚██████╔╝██║ ╚████║███████╗   ██║       ██║  ██║██║  ██╗
╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝   ╚═╝       ╚═╝  ╚═╝╚═╝  ╚═╝
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.Init(logLevel)
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "info", "设置日志等级 (Set log level) [trace|debug|info|warn|error|fatal|panic]")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  RootCmd,
		Headings: cc.HiGreen + cc.Underline,
		Commands: cc.Cyan + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Cyan + cc.Bold,
	})
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
