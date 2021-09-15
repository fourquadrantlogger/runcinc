package main

import (
	"github.com/sirupsen/logrus"
	"runcic/cmd"
)

func init() {
	//logrus.SetFormatter(&logrus.JSONFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//	PrettyPrint:     true,
	//})
	logrus.SetLevel(logrus.DebugLevel)
}
func main() {
	cmd.Execute()
}
