package fs

import (
	"github.com/sirupsen/logrus"
	"os"
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

var DefaultMounts = []MountConfig{
	devtmpfs, shm,
	devpts, ptmx,
	mqueue,
	proc,
	sys,
}

func Mount() (err error) {
	for i := 0; i < len(DefaultMounts); i++ {
		mc := DefaultMounts[i]
		if err := os.MkdirAll(mc.Target, 0o755); err != nil {
			continue
		}
		mc.Data = strings.Join(mc.Options, ",")
		mc.Options = nil
		err = syscall.Mount(mc.Source, mc.Target, mc.Fstype, 0, mc.Data)
		if err != nil {
			logrus.Errorf("unix.Mount %+v failed %s", mc, err.Error())
		} else {
			logrus.Infof("syscall.Mount(Source=%s, Target=%s, Fstype=%s,Data=%+v", mc.Source, mc.Target, mc.Fstype, mc.Data)
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
