package cic

import (
	"github.com/sirupsen/logrus"
	"runcic/cic/capabilities"
	"runcic/containerimage"
	"runcic/containerimage/common"
)

func mergeCap(CapAdd, CapDrop []string) {
	hashCap := make(map[string]bool)
	for _, c := range capabilities.DefaultCapabilities {
		hashCap[c] = true
	}
	for _, a := range CapAdd {
		hashCap[a] = true
	}
	for _, d := range CapDrop {
		delete(hashCap, d)
	}
	defaultCap := make([]string, 0)
	for c, _ := range hashCap {
		defaultCap = append(defaultCap, c)
	}
	capabilities.DefaultCapabilities = defaultCap
}
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
	mergeCap(cfg.CapAdd, cfg.CapDrop)
	run.Caps, err = capabilities.New(&capabilities.Capabilities{
		Bounding:    capabilities.DefaultCapabilities,
		Effective:   capabilities.DefaultCapabilities,
		Inheritable: capabilities.DefaultCapabilities,
		Permitted:   capabilities.DefaultCapabilities,
		Ambient:     capabilities.DefaultCapabilities,
	})
	if err != nil {
		logrus.WithField("err", err.Error()).Errorf("capabilities.New fail")
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
			logrus.WithField("image", run.Images[i].Image).Debug("spec image")
			imagespec := containerimage.Driver().Spec(run.Images[i].Image)
			if imagespec == nil {
				logrus.
					WithField("driver", containerimage.Driver().Name()).
					WithField("image", run.Images[i].Image).
					Info("image not found")
				pullimage(run.Images[i].Image, cfg.Authfile)
			}
		}
	}
	for i := 0; i < len(run.Images); i++ {
		imgi := containerimage.Driver().Spec(run.Images[i].Image)
		if imgi == nil {
			logrus.
				WithField("image", run.Images[i].Image).
				WithField("driver", containerimage.Driver().Name()).
				Warn("image spec is nil!")
			return
		}
		run.Images[i] = imgi
		logrus.
			WithField("image", run.Images[i].Image).
			WithField("driver", containerimage.Driver().Name()).
			WithField("imagespec", imgi).
			Debug()
	}

	//todo 创建之前，需要检测是否已存在
	run.mergeEnv(cfg.Env)
	run.mergeCmd()
	if run.Name == "" {
		run.ContainerID = newID()
		run.Name = newName()
	}

	if err = run.Create(); err != nil {
		logrus.Errorf("create cic by images %+v fail,error: %+v", run.ImageArray(), err.Error())
		return
	}

	if err = run.Start(); err != nil {
		logrus.Errorf("start image %+v %+v fail,error: %+v", run.ImageArray(), run.Command, err.Error())
		return
	}
	return
}
