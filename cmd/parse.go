package cmd

import (
	"errors"
	"regexp"
	"runcic/pkg/rflag"
	"strings"
)

func parse(args []string) (cmd, image, env []string, imageRoot, registrySecret, cicVolume, name string, copyParentEnv bool, err error) {
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
	imageIndex := imageIdx(args)
	var flags = make(map[string][]string)

	if imageIndex < 0 {

		_, unknownCmds := rflag.ParseFlag(args, []string{"env"})
		if len(unknownCmds) > 0 {
			image = strings.Split(unknownCmds[0], ",")
		} else {
			err = errors.New("run ? image")
			return
		}
		for i := 0; i < len(args); i++ {
			if args[i] == unknownCmds[0] {
				imageIndex = i
				break
			}
		}
	} else {
		cmd = args[imageIndex+1:]
		image = strings.Split(args[imageIndex], ",")
	}

	flags, _ = rflag.ParseFlag(args[:imageIndex], []string{"env"})
	if _, h := flags["env"]; h {
		env = flags["env"]
	}
	if _, h := flags["cicimage"]; h {
		imageRoot = flags["cicimage"][0]
	}
	if _, h := flags["registrysecret"]; h {
		registrySecret = flags["registrysecret"][0]
	}
	if _, h := flags["cicvolume"]; h {
		cicVolume = flags["cicvolume"][0]
	}
	if _, h := flags["name"]; h {
		name = flags["name"][0]
	}
	_, copyParentEnv = flags["copyenv"]
	return

}
