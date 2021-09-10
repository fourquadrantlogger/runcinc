package cic

import (
	"github.com/sirupsen/logrus"
	"runcic/containerimage"
)

func pullimage(img, authfile string) (pullerr error) {
	logrus.Infof("runcic imagedriver image pull --authfile=%s  %s", authfile, img)
	pullerr = containerimage.Driver().Pull(img, authfile)
	if pullerr != nil {

	} else {
		logrus.Infof("runcic imagedriver image pulled %s", img)
	}
	return
}
