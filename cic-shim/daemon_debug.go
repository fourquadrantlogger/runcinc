package cic_shim

import (
	"os"
	"syscall"
)

const daemonFlagName = "--daemon"

func initDaemonRuntime() {
	// 创建新回话
	_, err := syscall.Setsid()
	if err != nil {
		return
	}
	// 把标准输入输出指向null
	fd, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		return
	}
	_ = syscall.Dup2(int(fd.Fd()), int(os.Stdin.Fd()))
	_ = syscall.Dup2(int(fd.Fd()), int(os.Stdout.Fd()))
	_ = syscall.Dup2(int(fd.Fd()), int(os.Stderr.Fd()))
	if fd.Fd() > os.Stderr.Fd() {
		_ = fd.Close()
	}
}
func Daemon() (int, error) {
	isDaemon := false
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == daemonFlagName {
			isDaemon = true
		}
	}
	if isDaemon { // daemon process
		initDaemonRuntime()
		return 0, nil
	}
	procPath := os.Args[0]
	// 添加"--daemon"参数
	args := make([]string, 0, len(os.Args)+1)
	args = append(args, os.Args...)
	args = append(args, daemonFlagName)
	attr := &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}
	pid, err := syscall.ForkExec(procPath, args, attr)
	if err != nil {
		return -1, err
	}
	return pid, nil
}
