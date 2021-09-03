package cmd

import (
	"fmt"
	"os"
	"regexp"
	"runcic/pkg/rflag"
	"strings"
)

func parse(args []string) {
	var imageIdx = func(args []string) (imageIndex int) {
		imagePattern := `[A-Za-z0-9_\.\-/]+:[A-Za-z0-9_\.\-/]+`
		imageIndex = -1
		for i := 0; i < len(args); i++ {
			if match, _ := regexp.MatchString(imagePattern, args[i]); match {
				imageIndex = i
				break
			}
		}
		return
	}
	endIdx := imageIdx(args)
	if endIdx < 0 {
		endIdx = len(args) - 1
	} else {
		cmd = args[endIdx+1:]
	}
	var flags = make(map[string][]string)
	var unknownCmds []string
	flags, unknownCmds = rflag.ParseFlag(args[:endIdx-1], []string{"env"})
	if endIdx == len(args)-1 {
		fmt.Errorf("%+v", unknownCmds)
	} else {
		images = strings.Split(args[endIdx+1], ",")
	}

	if _, h := flags["cicimage"]; h {
		imageRoot = flags["cicimage"][0]
	}
	if _, h := flags["cicvolume"]; h {
		cicVolume = flags["cicvolume"][0]
	}
	_, copyParentEnv = flags["envcopy"]
	return

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
