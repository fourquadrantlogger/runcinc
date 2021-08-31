package cic

import (
	"github.com/sirupsen/logrus"
	"runcic/cic/fs"
)

func (r *Runcic) Start() (err error) {
	if err = r.cicvolume(); err != nil {
		return
	}
	if err = r.rootfspath(); err != nil {
		return
	}
	if err = r.mountoverlay(); err != nil {
		return
	}

	if r.ParentRootfs, err = realChroot(r.Roorfs()); err != nil {
		logrus.Errorf("chroot failed %s", err.Error())
		return
	}

	if err = fs.Mount(); err != nil {
		logrus.Errorf("fs mount failed %s", err.Error())
	}
	if err = fs.Link(); err != nil {
		logrus.Errorf("fs link failed %s", err.Error())
	}
	go func() {
		logrus.Infof("%+v %+v", r.Command, r.Envs)
		err = Execv(r.Command[0], r.Command[1:], r.Envs)
		if err != nil {
			logrus.Errorf("exec exited %v", err.Error())
		}
	}()

	return err
}
