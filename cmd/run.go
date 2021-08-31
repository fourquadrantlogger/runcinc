package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runcic/cic"
	"strings"
)

var (
	envs            []string
	name            string
	imagePullPolicy string = string(cic.ImagePullPolicyfNotPresent)
	imageRoot       string = "/image"
	cicVolume       string = ""
	ciccmd          string
)
var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run -e myenv:abc --name mycic myimage:latest --cmd `sh -c sleep 10h`",
	Long:  ``,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (vs []string, s cobra.ShellCompDirective) {
		return
	},
	Run: func(cmd *cobra.Command, args []string) {

		image := args
		cfg := cic.CicConfig{
			envs,
			strings.Split(ciccmd, " "),
			cic.ImagePullPolicy(imagePullPolicy),
			image,
			imageRoot,
			name,
			cicVolume,
		}
		log.Infof("runcic begin run:config %+v", cfg)
		cic.Run(cfg)
	},
}

func init() {
	flagSet := cmdRun.Flags()
	flagSet.SetInterspersed(false)
	flagSet.StringSliceVarP(&envs, "env", "e", []string{}, "--env VAR1=value1")
	flagSet.StringVar(&ciccmd, "cmd", "", "--cmd `sleep 10h`")
	flagSet.StringVar(&name, "name", "", "--name mycic")
	flagSet.StringVar(&cicVolume, "cicvolume", "", "--cicvolume /cicvolume")
	flagSet.StringVar(&imageRoot, "imageroot", "/image", "--imageroot /image")
	flagSet.StringVar(&imagePullPolicy, "imagePullPolicy", "", "--imagePullPolicy IfNotPresent")
}
