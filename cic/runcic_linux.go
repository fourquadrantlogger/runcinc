package cic

import (
	"fmt"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
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
	logrus.Infof("chrooting %s", path)
	if err := syscall.Chroot(path); err != nil {
		return fmt.Errorf("Error after fallback to chroot: %v", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("Error changing to new root after chroot: %v", err)
	}
	logrus.Infof("chroot success %s", path)
	return nil
}

//Execv
// https://github.com/opencontainers/runc/blob/master/libcontainer/system/linux.go
func Execv(cmd string, args []string, env []string) error {
	logrus.Infof("execv ing %s %+s %+v", cmd,args,env)
	name, err := exec.LookPath(cmd)
	if err != nil {
		logrus.Errorf("exec.LookPath(cmd) %s", err.Error())
		return err
	}

	return Exec(name, args, env)
}

func Exec(cmd string, args []string, env []string) error {
	for {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		logrus.Infof("exec ing %s %+s %+v", cmd,args,env)
		err := syscall.Exec(cmd, args, env)
		logrus.Errorf("syscall.Exec(cmd, args, env) %s", err.Error())
		if err != syscall.EINTR { //nolint:errorlint // unix errors are bare
			return err
		}
	}
}
