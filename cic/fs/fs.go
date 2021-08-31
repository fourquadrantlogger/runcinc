package fs

import (
	"github.com/sirupsen/logrus"
	"os"
	"runcic/utils"
	"strings"
	"syscall"
)

type MountConfig struct {
	Source  string
	Target  string
	Fstype  string
	Flags   uintptr
	Data    string
	Options []string
}
type LinkConfig struct {
	Source string
	Target string
}

var DefaultMounts = []MountConfig{
	devtmpfs,
	devpts,
	shm,
	mqueue,
	proc,
	sys,
}
var DefaultLinks = []LinkConfig{
	ptmx,
}

func Mount() (err error) {
	for i := 0; i < len(DefaultMounts); i++ {
		mc := DefaultMounts[i]
		if err := os.MkdirAll(mc.Target, 0o755); err != nil {
			logrus.Warnf("mkdir   %+v failed err:%s", mc.Target, err.Error())
		}
		mc.Data = strings.Join(mc.Options, ",")
		mc.Options = nil
		err = syscall.Mount(mc.Source, mc.Target, mc.Fstype, mc.Flags, mc.Data)
		if err != nil {
			logrus.Errorf("unix.Mount %+v failed %s", mc, err.Error())
		} else {
			logrus.Infof("syscall.Mount(Source=%s, Target=%s, Fstype=%s,Data=%+v", mc.Source, mc.Target, mc.Fstype, mc.Data)
		}
	}
	return
}
func Link() (err error) {
	for i := 0; i < len(DefaultLinks); i++ {
		mc := DefaultLinks[i]
		if utils.Exists(mc.Target) {
			err = os.Remove(mc.Target)
			if err != nil {
				logrus.Errorf("delete %+v failed %s", mc.Target, err.Error())
				return
			} else {
				logrus.Infof("syscall.link(Source=%s, Target=%s", mc.Source, mc.Target)
			}
		}

		err = os.Symlink(mc.Source, mc.Target)
		if err != nil {
			logrus.Errorf("link %+v failed %s", mc, err.Error())
		} else {
			logrus.Infof("syscall.link(Source=%s, Target=%s", mc.Source, mc.Target)
		}
	}
	return
}
func Umount() (err error) {
	for i := len(DefaultMounts) - 1; i >= 0; i-- {
		mc := DefaultMounts[i]
		err = syscall.Unmount(mc.Target, 0)
		if err != nil {
			logrus.Errorf("unix.UMount %+v failed %s", mc.Target, err.Error())
		} else {
			logrus.Infof("syscall.UMount(Target=%s", mc.Target)
		}
	}
	return
}
