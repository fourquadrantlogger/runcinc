package cmd

import (
	"os"
	"runcic/cic"
)

var cfg = cic.CicConfig{
	[]string{},
	false,
	[]string{},
	cic.ImagePullPolicyfNotPresent,
	[]string{},
	"/image",
	"",
	"",
}

// Execute executes the root command.
func Execute() {
	if len(os.Args) > 2 {

		parse(os.Args[2:])
		switch os.Args[1] {
		case "runin":
			cmdRun()
		case "run":
			cmdWait()
		}
	} else {
		cmdHelp()
	}
}
