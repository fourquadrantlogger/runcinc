package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "runcic",
		Short: "一个可以在标准父容器（cic）内run 子容器image的工具",
		Long:  `runcic依赖podman获取container image的layer信息`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(cmdRun)
}
