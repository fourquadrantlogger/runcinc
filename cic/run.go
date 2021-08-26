package cic

import (
	"github.com/sirupsen/logrus"
	"runcic/containerimage"
)

func Run(cfg CicConfig) (err error) {

	run := &Runcic{
		Envs:    cfg.Env,
		Name:    cfg.Name,
		Command: cfg.Cmd,
		Image: &containerimage.Image{
			Image: cfg.Image,
		},
		ImagePullPolicy: cfg.ImagePullPolicy,
	}
	switch run.ImagePullPolicy {
	case imagePullPolicyAlways:
		containerimage.Driver().Pull(run.Image.Image)
	case ImagePullPolicyfNotPresent:
		imagespec := containerimage.Driver().Spec(run.Image.Image)
		if imagespec == nil {
			containerimage.Driver().Pull(run.Image.Image)
		}
	}
	run.Image = containerimage.Driver().Spec(run.Image.Image)
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
		logrus.Errorf("start cic by image %s fail,error: %s", run.Image.Image, err.Error())
		return
	}
	return
}
