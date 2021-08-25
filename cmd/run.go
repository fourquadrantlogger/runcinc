package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	envs []string
)
var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run image -e myenv:abc sh -c sleep 10h",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, kv_ := range envs {
			splitIndex := strings.IndexAny(kv_, "=")
			if splitIndex > 0 {
				fmt.Println("export ", kv_[:splitIndex], "=", kv_[splitIndex+1:])
			}
		}
		fmt.Println("runcic run " + strings.Join(args, " "))
	},
}

func init() {
	cmdRun.Flags().StringSliceVarP(&envs, "env", "e", []string{"qf:fw"}, "wf")
}
