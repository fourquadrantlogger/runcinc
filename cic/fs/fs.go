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
	data    string
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
			return err
		}
		err = syscall.Mount(mc.Source, mc.Target, mc.Fstype, 0, strings.Join(mc.Options, ","))
		if err != nil {
			logrus.Errorf("unix.Mount %+v failed %s", mc, err.Error())
		} else {
			logrus.Infof("syscall.Mount(Source=%s, Target=%s, Fstype=%s,Options=%+v", mc.Source, mc.Target, mc.Fstype, mc.Options)
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
