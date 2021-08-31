package cic

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"runcic/containerimage/common"
	"runcic/utils"
	"strings"
	"time"
)

type Runcic struct {
	ParentRootfs    *os.File
	cancel          context.CancelFunc
	Name            string
	CicVolume       string
	ContainerID     string
	Images          []*common.Image
	Command         []string
	Envs            []string
	Started         time.Time
	ImagePullPolicy ImagePullPolicy
}

func (r *Runcic) ImageArray() (imgs []string) {
	for i := 0; i < len(r.Images); i++ {
		imgs = append(imgs, r.Images[i].Image)
	}
	return
}
func (r *Runcic) Roorfs() (path string) {
	path = OverlayRoot + string(os.PathSeparator) + r.Name
	return
}

func (r *Runcic) mountops() string {
	lower := make([]string, 0)
	for i := 0; i < len(r.Images); i++ {
		lower = append(lower, r.Images[i].Lower...)
	}
	mountops := strings.Join([]string{
		"lowerdir=" + strings.Join(lower, ":"),
		"upperdir=" + r.CicVolume + string(os.PathSeparator) + "up",
		"workdir=" + r.CicVolume + string(os.PathSeparator) + "work",
	}, ",")
	return mountops
}

func (r *Runcic) rootfspath() (err error) {
	work, worke := os.Stat(r.Roorfs())
	if worke != nil {
		utils.Mkdirp(r.Roorfs())
	} else {
		if !work.IsDir() {
			err = errors.New("rootfs should be dir!" + r.Roorfs())
			logrus.Warnf(err.Error())
		}
	}
	logrus.Infof("rootfs ok,%s", r.Roorfs())
	return
}
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
	logrus.Infof("cicvolume updir ok,%s", r.CicVolume+"/"+"up")

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

	logrus.Infof("cicvolume workdir ok,%s", r.CicVolume+"/"+"work")
	return
}
