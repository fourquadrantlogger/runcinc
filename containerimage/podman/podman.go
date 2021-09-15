package podman

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runcic/containerimage/common"
	"strings"
	"time"
)

type Podman struct {
	Root string
	Auth string
}

func (c *Podman) Name() string {
	return "podman"
}
func (c *Podman) Spec(image string) (img *common.Image) {
	cmds := []string{
		"podman",
		"--root", c.Root,
		"--storage-driver=overlay2",
		"inspect", image}
	speccmd := exec.Command(cmds[0], cmds[1:]...)
	speccmd.Stderr = os.Stderr
	var out bytes.Buffer
	speccmd.Stdout = &out
	logrus.WithField("cmds", strings.Join(cmds, " ")).Info()
	err := speccmd.Run()

	if err != nil {
		logrus.WithField("err", err.Error()).Error("podman image inspect failed")
		return
	}

	var images = make([]podmanImageInspect, 0)
	err = json.Unmarshal(out.Bytes(), &images)
	if err != nil {
		logrus.WithField("image", image).WithField("err", err.Error()).Error("unmarshal podman inspect image failed")
		logrus.Error(out.String())
		return
	}
	if len(images) > 0 {
		img = &common.Image{}
		if len(images[0].RepoTags) > 0 {
			img.Image = images[0].RepoTags[0]
		}
		img.Env = images[0].Config.Env
		img.Cmd = images[0].Config.Cmd
		lower := strings.Split(strings.TrimSpace(images[0].GraphDriver.Data.LowerDir), ":")
		for i := 0; i < len(lower); i++ {
			if strings.TrimSpace(lower[i]) != "" {
				img.Lower = append(img.Lower, lower[i])
			}
		}
		img.Lower = append(img.Lower, strings.TrimSpace(images[0].GraphDriver.Data.UpperDir))
	}
	return
}
func (c *Podman) Pull(image, authfile string) (err error) {
	cmds := []string{
		"podman",
		"--storage-driver=overlay2",
		"--root=" + c.Root,
		"image", "pull",
	}
	if authfile != "" {
		cmds = append(cmds, "--authfile", authfile)
	}
	cmds = append(cmds, image)

	pullcmd := exec.Command(cmds[0], cmds[1:]...)
	pullcmd.Stdout = os.Stdout
	pullcmd.Stderr = os.Stderr
	logrus.WithField("cmds", strings.Join(cmds, " ")).Info()
	err = pullcmd.Run()
	if err != nil {
		logrus.WithField("err", err.Error()).Error("podman image pull failed")
		return
	}
	return
}

type podmanImageInspect struct {
	Id          string    `json:"Id"`
	Digest      string    `json:"Digest"`
	RepoTags    []string  `json:"RepoTags"`
	RepoDigests []string  `json:"RepoDigests"`
	Parent      string    `json:"Parent"`
	Comment     string    `json:"Comment"`
	Created     time.Time `json:"Created"`
	Config      struct {
		ExposedPorts struct {
			Tcp struct {
			}
		} `json:"ExposedPorts"`
		Env        []string `json:"Env"`
		Entrypoint []string `json:"Entrypoint"`
		Cmd        []string `json:"Cmd"`
		Volumes    struct {
			Data struct {
			}
		} `json:"Volumes"`
		WorkingDir string `json:"WorkingDir"`
	} `json:"Config"`
	Version      string `json:"Version"`
	Author       string `json:"Author"`
	Architecture string `json:"Architecture"`
	Os           string `json:"Os"`
	Size         int    `json:"Size"`
	VirtualSize  int    `json:"VirtualSize"`
	GraphDriver  struct {
		Name string `json:"Name"`
		Data struct {
			LowerDir string `json:"LowerDir"`
			UpperDir string `json:"UpperDir"`
			WorkDir  string `json:"WorkDir"`
		} `json:"Data"`
	} `json:"GraphDriver"`
	RootFS struct {
		Type   string   `json:"Type"`
		Layers []string `json:"Layers"`
	} `json:"RootFS"`
	Labels      interface{} `json:"Labels"`
	Annotations struct {
	} `json:"Annotations"`
	ManifestType string `json:"ManifestType"`
	User         string `json:"User"`
	History      []struct {
		Created    time.Time `json:"created"`
		CreatedBy  string    `json:"created_by"`
		EmptyLayer bool      `json:"empty_layer,omitempty"`
	} `json:"History"`
	NamesHistory []string `json:"NamesHistory"`
}
