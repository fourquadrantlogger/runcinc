package cmd

import (
	log "github.com/sirupsen/logrus"
	"runcic/cic"
)

func cmdRun() {

	log.Infof("runcic run:config %+v", cfg)
	cic.Run(cfg)
}
