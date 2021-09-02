package cic

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func Execc(ctx context.Context, cmd string, args []string, env []string) (err error) {
	name, err := exec.LookPath(cmd)
	if err != nil {
		logrus.Infof("exec.LookPath %s not found,error %v", err.Error())
		return err
	}
	c := exec.CommandContext(ctx, name, args...)
	c.Env = env
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	logrus.Infof("exec %s %+v env[%+v]", cmd, args, env)
	return c.Run()
}
