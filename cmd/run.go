package cmd

import (
	"runcic/cic"
	cic_shim "runcic/cic-shim"
)

var cmdWait = func() {

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

	cic_shim.Wait(cfg)
}
