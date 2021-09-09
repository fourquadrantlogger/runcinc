package cic

import (
	"github.com/sirupsen/logrus"
	"runcic/containerimage"
	"runcic/containerimage/common"
)

func Run(cfg CicConfig) (err error) {
	run := &Runcic{
		Volume:          cfg.Volume,
		CopyEnv:         cfg.CopyParentEnv,
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
	var pullimage = func(img, authfile string) (pullerr error) {
		logrus.Infof("runcic imagedriver image pull --authfile=%s  %s", authfile, img)
		pullerr = containerimage.Driver().Pull(img, authfile)
		if pullerr != nil {

		} else {
			logrus.Infof("runcic imagedriver image pulled %s", img)
		}
		return
	}
	switch run.ImagePullPolicy {
	case imagePullPolicyAlways:
		for i := 0; i < len(run.Images); i++ {
			err = pullimage(run.Images[i].Image, cfg.Authfile)
			if err != nil {
				return
			}
		}
	case ImagePullPolicyfNotPresent:
		fallthrough
	default:
		for i := 0; i < len(run.Images); i++ {
			logrus.Infof("runcic imagedriver spec image %s", run.Images[i].Image)
			imagespec := containerimage.Driver().Spec(run.Images[i].Image)
			if imagespec == nil {
				logrus.Warnf("runcic imagedriver not found image %s", run.Images[i].Image)
				pullimage(run.Images[i].Image, cfg.Authfile)
			}
		}
	}
	for i := 0; i < len(run.Images); i++ {
		imgi := containerimage.Driver().Spec(run.Images[i].Image)
		run.Images[i] = imgi
		if imgi == nil {
			logrus.Errorf("runcic imagedriver spec image is nil,your image=%s", run.Images[i].Image)
			return
		}
		logrus.Infof("runcic imagedriver spec image %+v", imgi)
	}

	//todo 创建之前，需要检测是否已存在
	run.mergeEnv(cfg.Env)
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
