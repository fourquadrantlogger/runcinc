package cic

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"runcic/cic/fs"
	"syscall"
)

func (r *Runcic) postStop(s os.Signal, oldpath *os.File) {
	logrus.Infof("recv signal %+v", s)
	logrus.Info("sending cic signal")
	r.cancel()
	logrus.Info("cic exited")
	fs.Umount()
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

func (r *Runcic) Stop() (err error) {
	return
}

func (r *Runcic) WaitSignal(f func(sig os.Signal, oldpath *os.File)) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	s := <-sigs
	f(s, r.ParentRootfs)
}
