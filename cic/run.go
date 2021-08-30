package cic

import (
	"github.com/sirupsen/logrus"
	"runcic/containerimage"
	"runcic/containerimage/common"
	"runcic/containerimage/podman"
)

func Run(cfg CicConfig) (err error) {

	run := &Runcic{
		Envs:      cfg.Env,
		Name:      cfg.Name,
		Command:   cfg.Cmd,
		CicVolume: cfg.CicVolume,
		Image: &common.Image{
			Image: cfg.Image,
		},
		ImagePullPolicy: cfg.ImagePullPolicy,
	}
	containerimage.SetDriver(&podman.Podman{
		Root: cfg.ImageRoot,
	})
	var pullimage = func() {
		logrus.Infof("runcic imagedriver pulling image %s", run.Image.Image)
		containerimage.Driver().Pull(run.Image.Image)
		logrus.Infof("runcic imagedriver pulled image %s", run.Image.Image)
	}
	switch run.ImagePullPolicy {
	case imagePullPolicyAlways:
		pullimage()
	default:
		logrus.Infof("runcic imagedriver spec image %s", run.Image.Image)
		imagespec := containerimage.Driver().Spec(run.Image.Image)
		if imagespec == nil {
			logrus.Warnf("runcic imagedriver not found image %s", run.Image.Image)
			pullimage()
		}
	}
	run.Image = containerimage.Driver().Spec(run.Image.Image)
	if run.Image == nil {
		return
	}
	logrus.Infof("runcic imagedriver spec image %+v", run.Image)
	//todo 创建之前，需要检测是否已存在
	if run.Name == "" {
		if err = run.Create(); err != nil {
			logrus.Errorf("create cic by image %s fail,error: %s", run.Image.Image, err.Error())
			return
		}
	} else {
		//todo 已存在
	}

	if err = run.Start(); err != nil {
		logrus.Errorf("start image %s %+v fail,error: %s", run.Image.Image, run.Command, err.Error())
		return
	}
	return
}
