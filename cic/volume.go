package cic

import (
	"github.com/sirupsen/logrus"
	"runcic/utils"
	"strings"
	"syscall"
)

func (r *Runcic) mountbindvolume() (err error) {
	for i := 0; i < len(r.Volume); i++ {
		v := strings.Split(r.Volume[i], ":")
		if len(v) >= 2 {
			source, target := v[0], r.Roorfs()+v[1]
			utils.Mkdirp(target)
			err = syscall.Mount(source, target, "bind", syscall.MS_BIND|syscall.MS_REC, "")
			if err != nil {
				logrus.Errorf("mount bind %s:%s err %+v", source, target, err.Error())
				return err
			} else {
				logrus.Infof("mount bind %s %s", source, target)
			}
		} else {
			logrus.Errorf("error volume %s", r.Volume[i])
		}
	}
	return
}
