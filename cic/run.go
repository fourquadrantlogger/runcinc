package cic

import (
	"github.com/sirupsen/logrus"
	"runcic/containerimage"
	"runcic/containerimage/common"
)

func Run(cfg CicConfig) (err error) {
	run := &Runcic{
		Envs:            cfg.Env,
		Name:            cfg.Name,
		Command:         cfg.Cmd,
		CicVolume:       cfg.CicVolume,
		ImagePullPolicy: cfg.ImagePullPolicy,
	}
	for i := 0; i < len(cfg.Images); i++ {
		run.Images = append(run.Images, &common.Image{
			Image: cfg.Images[i],
		})
	}
	var pullimage = func(img string) (pullerr error) {
		logrus.Infof("runcic imagedriver pulling image %s", img)
		pullerr = containerimage.Driver().Pull(img)
		if pullerr != nil {

		} else {
			logrus.Infof("runcic imagedriver pulled image %s", img)
		}
		return
	}
	switch run.ImagePullPolicy {
	case imagePullPolicyAlways:
		for i := 0; i < len(run.Images); i++ {
			err = pullimage(run.Images[i].Image)
			if err != nil {
				return
			}
		}
	default:
		for i := 0; i < len(run.Images); i++ {
			logrus.Infof("runcic imagedriver spec image %s", run.Images[i].Image)
			imagespec := containerimage.Driver().Spec(run.Images[i].Image)
			if imagespec == nil {
				logrus.Warnf("runcic imagedriver not found image %s", run.Images[i].Image)
				pullimage(run.Images[i].Image)
			}

		}
	}
	for i := 0; i < len(run.Images); i++ {
		imgi := containerimage.Driver().Spec(run.Images[i].Image)
		run.Images[i] = imgi
		if imgi == nil {
			return
		}
		logrus.Infof("runcic imagedriver spec image %+v", imgi)
	}

	//todo 创建之前，需要检测是否已存在
	run.mergeEnv()
	run.mergeCmd()
	if run.Name == "" {
		run.ContainerID = newID()
		run.Name = newName()
	}

	if err = run.Create(); err != nil {
		logrus.Errorf("create cic by images %+v fail,error: %s", run.ImageArray(), err.Error())
		return
	}

	if err = run.Start(); err != nil {
		logrus.Errorf("start image %+v %+v fail,error: %s", run.ImageArray(), run.Command, err.Error())
		return
	}
	return
}
