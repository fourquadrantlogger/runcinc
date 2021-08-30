package cic

import (
	"os"
	"os/signal"
	"syscall"
)

func (r *Runcic) Stop() (err error) {
	return
}

func (r *Runcic) WaitSignal(f func(sig os.Signal, oldpath *os.File)) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	s := <-sigs
	f(s, r.ParentRootfs)
}
