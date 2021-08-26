package cic

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
)

func (r *Runcic) cicvolume() (err error) {
	if r.CicVolume == "" {
		err = errors.New("cicvolume not set!")
		logrus.Errorf(err.Error())
		return err
	}
	stat, fe := os.Stat(r.CicVolume)
	if fe != nil {
		logrus.Errorf(fe.Error())
		return fe
	}
	if !stat.IsDir() {
		err = errors.New("cicvolume should be dir!,fix your cicvolume: " + r.CicVolume)
		logrus.Errorf(err.Error())
		return
	}

	up, upe := os.Stat(r.CicVolume + "/" + "up")
	if upe != nil {
		err = os.Mkdir(r.CicVolume+"/"+"up", os.ModeDir)
		if err != nil {
			logrus.Errorf("mkdir cicvolume/up dir fail,error: %s", err.Error())
			return err
		}
	} else {
		if !up.IsDir() {
			err = errors.New("cicvolume/up should be dir!" + r.CicVolume)
			logrus.Warnf(err.Error())
		}
	}

	work, worke := os.Stat(r.CicVolume + "/" + "work")
	if worke != nil {
		err = os.Mkdir(r.CicVolume+"/"+"work", os.ModeDir)
		if err != nil {
			logrus.Errorf("mkdir cicvolume/work dir fail,error: %s", err.Error())
			return err
		}
	} else {
		if !work.IsDir() {
			err = errors.New("cicvolume/work should be dir!" + r.CicVolume)
			logrus.Warnf(err.Error())
		}
	}
	return
}

func (r *Runcic) Start() (err error) {
	if err = r.cicvolume(); err != nil {
		return
	}
	if err = r.mountoverlay(); err != nil {
		return
	}
	if err = realChroot(r.Roorfs()); err != nil {
		return
	}
	err = Execv(r.Command[0], r.Command[1:], r.Envs)
	return
}
