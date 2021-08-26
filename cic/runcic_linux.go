package cic

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"os/exec"
)

func (r *Runcic) mountoverlay() (err error) {
	mountops := r.mountops()
	err = unix.Mount("overlay", r.Roorfs(), "overlay", 0, mountops)
	if err != nil {
		logrus.Errorf("mount overlay fail,errors %s", err.Error())
	}
	return err
}
func realChroot(path string) error {
	if err := unix.Chroot(path); err != nil {
		return fmt.Errorf("Error after fallback to chroot: %v", err)
	}
	if err := unix.Chdir("/"); err != nil {
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
		err := unix.Exec(cmd, args, env)
		if err != unix.EINTR { //nolint:errorlint // unix errors are bare
			return err
		}
	}
}
