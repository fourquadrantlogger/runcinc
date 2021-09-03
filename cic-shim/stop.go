package cic_shim

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"runcic/cic/fs"
	"syscall"
)

func (r *RuncicShim) postStop(s os.Signal, oldpath *os.File) {
	logrus.Infof("recv signal %+v", s)
	logrus.Info("sending cic signal")
	r.cancel()
	logrus.Info("cic exited")
	fs.Umount("")
	err := oldpath.Chdir()
	if err != nil {
		logrus.Errorf("chdir() err: %v", err)
		return
	}
	err = syscall.Chroot(".")
	if err != nil {
		logrus.Errorf("chroot back err: %v", err)
		return
	} else {
		logrus.Infof("chroot back")
	}
	err = syscall.Unmount(r.Roorfs(), 0)
	if err != nil {
		logrus.Errorf("umount overlay failed %s,err: %v", r.Roorfs(), err.Error())
		return
	} else {
		logrus.Infof("umount overlay %v", r.Roorfs())
	}
	logrus.Infof("runcic exit")
	return
}

//https://github.com/containerd/containerd/blob/main/cmd/containerd-shim/shim_linux.go
func (r *RuncicShim) WaitSignal(f func(sig os.Signal, oldpath *os.File)) {
	signals := make(chan os.Signal, 32)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT, unix.SIGCHLD, unix.SIGPIPE)
	s := <-signals
	f(s, r.ParentRootfs)
}
