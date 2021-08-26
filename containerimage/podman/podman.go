package podman

import (
	"encoding/json"
	"fmt"
	"github.com/bitfield/script"
	log "github.com/sirupsen/logrus"
	"runcic/containerimage"
	"strings"
	"time"
)

func init() {
	containerimage.SetDriver(&Podman{})
}

type Podman struct {
}

func (c *Podman) Spec(image string) (img *containerimage.Image) {
	result, err := script.Exec(fmt.Sprintf("podman image inspect %s", image)).String()
	if err != nil {
		log.Errorf("exec failed: %v", err.Error())
		log.Errorf(result)
		return
	}
	var images = make([]podmanImageInspect, 0)
	err = json.Unmarshal([]byte(result), &images)
	if err != nil {
		log.Errorf("unmarshal podman inspect %s json failed: %v", image, err.Error())
		log.Errorf(result)
		return
	}
	if len(images) > 0 {
		img = &containerimage.Image{}
		if len(images[0].RepoTags) > 0 {
			img.Image = images[0].RepoTags[0]
		}
		img.Env = images[0].Config.Env
		img.Cmd = images[0].Config.Cmd
		img.Lower = strings.Split(images[0].GraphDriver.Data.LowerDir, ":")
	}
	return
}
func (c *Podman) Pull(image string) {
	script.Exec(fmt.Sprintf("docker image pull %s", image)).Stdout()
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
			} `json:"6379/tcp"`
		} `json:"ExposedPorts"`
		Env        []string `json:"Env"`
		Entrypoint []string `json:"Entrypoint"`
		Cmd        []string `json:"Cmd"`
		Volumes    struct {
			Data struct {
			} `json:"/data"`
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
