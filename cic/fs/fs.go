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
	devpts, ptmx,
	mqueue,
	proc,
	sys,
	devtmpfs, shm,
}

func Mount() (err error) {
	for _, mc := range DefaultMounts {
		if err := os.MkdirAll(mc.Target, 0o755); err != nil {
			return err
		}
		logrus.Infof("syscall.Mount(Source=%s, Target=%s, Fstype=%s,Options=%+v", mc.Source, mc.Target, mc.Fstype, mc.Options)
		err = syscall.Mount(mc.Source, mc.Target, mc.Fstype, 0, strings.Join(mc.Options, ","))
		if err != nil {
			logrus.Errorf("unix.Mount %+v failed %s", mc, err.Error())
		}
	}
	return
}
