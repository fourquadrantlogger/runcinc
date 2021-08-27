package cic

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"syscall"
)

func (r *Runcic) mountoverlay() (err error) {
	mountops := r.mountops()
	err = syscall.Mount("overlay", r.Roorfs(), "overlay", 0, mountops)
	logrus.Infof("mount overlay overlay -o %s %s ", mountops, r.Roorfs())
	if err != nil {
		logrus.Errorf("mount overlay fail,errors %s", err.Error())
		return
	}

	return err
}
func realChroot(path string) error {
	if err := syscall.Chroot(path); err != nil {
		return fmt.Errorf("Error after fallback to chroot: %v", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("Error changing to new root after chroot: %v", err)
	}
	return nil
}

//Execv
// https://github.com/opencontainers/runc/blob/master/libcontainer/system/linux.go
func Execv(cmd string, args []string, env []string) error {
	name, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}

	return Exec(name, args, env)
}

func Exec(cmd string, args []string, env []string) error {
	for {
		err := syscall.Exec(cmd, args, env)
		if err != syscall.EINTR { //nolint:errorlint // unix errors are bare
			return err
		}
	}
}
