package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "runcic",
		Short: "一个可以在容器内运行容器image的工具",
		Long:  `runcic,一个可以在容器内运行容器image的工具`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(cmdRun)
}
