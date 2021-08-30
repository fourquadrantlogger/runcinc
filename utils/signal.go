package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal(f func(sig os.Signal)) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	s := <-sigs
	f(s)
}
