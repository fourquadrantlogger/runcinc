package cmd

import (
	"github.com/spf13/cobra"
	"runcic/cic"
	cic_shim "runcic/cic-shim"
)

var cmdWait = &cobra.Command{
	Use:   "run",
	Short: "run -e myenv:abc --name mycic myimage:latest --cmd `sh -c sleep 10h`",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		cfg := cic.CicConfig{
			envs,
			copyParentEnv,
			args,
			cic.ImagePullPolicy(imagePullPolicy),
			images,
			imageRoot,
			name,
			cicVolume,
		}

		cic_shim.Wait(cfg)
	},
}
