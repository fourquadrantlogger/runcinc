package cmd

import (
	"github.com/sirupsen/logrus"
	"os"
	"runcic/cic"
)

var cfg = cic.CicConfig{
	[]string{},
	[]string{},
	[]string{},
	[]string{},
	false,
	[]string{},
	cic.ImagePullPolicyfNotPresent,
	[]string{},
	"/image",
	"",
	"",
	"",
}

// Execute executes the root command.
func Execute() {
	if len(os.Args) > 2 {
		var imageroot, cicvolume, name string
		logrus.Info(os.Args)
		cfg.Cmd, cfg.Images, cfg.Env, cfg.CapAdd, cfg.CapDrop, cfg.Volume,
			imageroot, cfg.Authfile, cicvolume, name,
			cfg.CopyParentEnv, _ = parse(os.Args[2:])
		if imageroot != "" {
			cfg.ImageRoot = imageroot
		}
		if cicvolume != "" {
			cfg.CicVolume = cicvolume
		}
		if name != "" {
			cfg.Name = name
		}
		switch os.Args[1] {
		case "runin":
			cmdRun()
		case "run":
			cmdWait()
		}
	} else {
		cmdHelp()
	}
}
