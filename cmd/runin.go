package cmd

import (
	log "github.com/sirupsen/logrus"
	"runcic/cic"
)

var (
	envs            []string
	copyParentEnv   bool
	name            string
	imagePullPolicy string //=string(cic.ImagePullPolicyfNotPresent)
	imageRoot       string = "/image"
	cicVolume       string = ""
	images          []string
	cmd             []string
)

var cmdRun = func() {
	cfg := cic.CicConfig{
		envs,
		copyParentEnv,
		cmd,
		cic.ImagePullPolicy(imagePullPolicy),
		images,
		imageRoot,
		name,
		cicVolume,
	}
	log.Infof("runcic run:config %+v", cfg)
	cic.Run(cfg)
}
