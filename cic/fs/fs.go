package fs

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"strings"
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
		err = unix.Mount(mc.Source, mc.Target, mc.Fstype, 0, strings.Join(mc.Options, ","))
		if err != nil {
			logrus.Errorf("unix.Mount %+v failed %s", mc, err.Error())
		}
	}
	return
}
