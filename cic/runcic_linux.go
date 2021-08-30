package cic

import (
	"fmt"
	"os"
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
	name, err := exec.LookPath(cmd)
	if err != nil {
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
		logrus.Infof("exec ing %s %+s %+v", cmd, args, env)
		err := syscall.Exec(cmd, args, env)
		if err != syscall.EINTR { //nolint:errorlint // unix errors are bare
			return err
		}
	}
}

func Execc(cmd string, args []string, env []string) (err error) {
	name, err := exec.LookPath(cmd)
	if err != nil {
		logrus.Infof("exec.LookPath %s not found,error %v", err.Error())
		return err
	}
	c := exec.Command(name, args...)
	c.Env = env
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	logrus.Infof("exec %s %+v env[%+v]")
	c.Start()
	return
}
