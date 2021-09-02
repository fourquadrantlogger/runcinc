package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"runcic/cic"
)

var (
	envs            []string
	name            string
	imagePullPolicy string = string(cic.ImagePullPolicyfNotPresent)
	imageRoot       string = "/image"
	cicVolume       string = ""
	images          []string
)
var cmdRun = &cobra.Command{
	Use:   "runin",
	Short: "runin -e myenv:abc --name mycic myimage:latest --cmd `sh -c sleep 10h`",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := cic.CicConfig{
			envs,
			args,
			cic.ImagePullPolicy(imagePullPolicy),
			images,
			imageRoot,
			name,
			cicVolume,
		}
		log.Infof("runcic run:config %+v", cfg)
		cic.Run(cfg)
	},
}

func init() {
	//cmdRun.Flags() the same as cmdWait.Flags()
	flagsets := []*pflag.FlagSet{cmdRun.Flags(), cmdWait.Flags()}
	for i := 0; i < len(flagsets); i++ {
		flagsets[i].SetInterspersed(false)
		flagsets[i].StringSliceVarP(&envs, "env", "e", []string{}, "--env VAR1=value1")
		flagsets[i].StringSliceVarP(&images, "image", "i", []string{}, "--image ubuntu:latest")
		flagsets[i].StringVar(&name, "name", "", "--name mycic")
		flagsets[i].StringVar(&cicVolume, "cicvolume", "", "--cicvolume /cicvolume")
		flagsets[i].StringVar(&imageRoot, "imageroot", "/image", "--imageroot /image")
		flagsets[i].StringVar(&imagePullPolicy, "imagePullPolicy", "", "--imagePullPolicy IfNotPresent")
	}

}
