package cmd

import (
	"gitee.com/xpigpig/vblog/cmd/start"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vblogApi",
	Short: "vblog后端项目",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// 需要执行命令的逻辑
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(start.Cmd)
}
