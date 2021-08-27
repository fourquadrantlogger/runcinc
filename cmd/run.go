package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runcic/cic"
)

var (
	envs            []string
	name            string
	imagePullPolicy string = string(cic.ImagePullPolicyfNotPresent)
	cicVolume       string = ""
)
var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run image -e myenv:abc --name mycic sh -c sleep 10h",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		image := args[0]
		cmds := args[1:]
		cfg := cic.CicConfig{
			envs,
			cmds,
			cic.ImagePullPolicy(imagePullPolicy),
			image,
			name,
			cicVolume,
		}
		log.Infof("runcic begin run:config %+v", cfg)
		cic.Run(cfg)
	},
}

func init() {
	flagSet := cmdRun.Flags()
	flagSet.StringSliceVarP(&envs, "env", "e", []string{}, "--env VAR1=value1")
	flagSet.StringVar(&name, "name", "", "--name mycic")
	flagSet.StringVar(&cicVolume, "cicvolume", "", "--cicvolume /cicvolume")
	flagSet.StringVar(&imagePullPolicy, "imagePullPolicy", "", "--imagePullPolicy IfNotPresent")
}
