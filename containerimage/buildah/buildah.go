package docker

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runcic/containerimage/common"
	"time"
)

type Docker struct {
}

func (c *Docker) Spec(image string) (img *common.Image) {

	return
}
func (c *Docker) Pull(image string) (err error) {
	log.Infof("buildah image  start pull %s", image)
	pullcmd := exec.Command("buildah", "pull", image)
	pullcmd.Stdout = os.Stdout
	err = pullcmd.Run()
	if err != nil {
		log.Errorf("buildah pull failed: %v", err.Error())
		return
	}
	return
}

type dockerImageInspect struct {
	Type             string `json:"Type"`
	FromImage        string `json:"FromImage"`
	FromImageID      string `json:"FromImageID"`
	FromImageDigest  string `json:"FromImageDigest"`
	Config           string `json:"Config"`
	Manifest         string `json:"Manifest"`
	Container        string `json:"Container"`
	ContainerID      string `json:"ContainerID"`
	MountPoint       string `json:"MountPoint"`
	ProcessLabel     string `json:"ProcessLabel"`
	MountLabel       string `json:"MountLabel"`
	ImageAnnotations struct {
	} `json:"ImageAnnotations"`
	ImageCreatedBy string `json:"ImageCreatedBy"`
	OCIv1          struct {
		Created      time.Time `json:"created"`
		Architecture string    `json:"architecture"`
		Os           string    `json:"os"`
		Config       struct {
			Env []string `json:"Env"`
			Cmd []string `json:"Cmd"`
		} `json:"config"`
		Rootfs struct {
			Type    string   `json:"type"`
			DiffIds []string `json:"diff_ids"`
		} `json:"rootfs"`
		History []struct {
			Created    time.Time `json:"created"`
			CreatedBy  string    `json:"created_by"`
			EmptyLayer bool      `json:"empty_layer,omitempty"`
		} `json:"history"`
	} `json:"OCIv1"`
	Docker struct {
		Created         time.Time `json:"created"`
		Container       string    `json:"container"`
		ContainerConfig struct {
			Hostname     string      `json:"Hostname"`
			Domainname   string      `json:"Domainname"`
			User         string      `json:"User"`
			AttachStdin  bool        `json:"AttachStdin"`
			AttachStdout bool        `json:"AttachStdout"`
			AttachStderr bool        `json:"AttachStderr"`
			Tty          bool        `json:"Tty"`
			OpenStdin    bool        `json:"OpenStdin"`
			StdinOnce    bool        `json:"StdinOnce"`
			Env          []string    `json:"Env"`
			Cmd          []string    `json:"Cmd"`
			ArgsEscaped  bool        `json:"ArgsEscaped"`
			Image        string      `json:"Image"`
			Volumes      interface{} `json:"Volumes"`
			WorkingDir   string      `json:"WorkingDir"`
			Entrypoint   interface{} `json:"Entrypoint"`
			OnBuild      interface{} `json:"OnBuild"`
			Labels       interface{} `json:"Labels"`
		} `json:"container_config"`
		Config struct {
			Hostname     string      `json:"Hostname"`
			Domainname   string      `json:"Domainname"`
			User         string      `json:"User"`
			AttachStdin  bool        `json:"AttachStdin"`
			AttachStdout bool        `json:"AttachStdout"`
			AttachStderr bool        `json:"AttachStderr"`
			Tty          bool        `json:"Tty"`
			OpenStdin    bool        `json:"OpenStdin"`
			StdinOnce    bool        `json:"StdinOnce"`
			Env          []string    `json:"Env"`
			Cmd          []string    `json:"Cmd"`
			ArgsEscaped  bool        `json:"ArgsEscaped"`
			Image        string      `json:"Image"`
			Volumes      interface{} `json:"Volumes"`
			WorkingDir   string      `json:"WorkingDir"`
			Entrypoint   interface{} `json:"Entrypoint"`
			OnBuild      interface{} `json:"OnBuild"`
			Labels       interface{} `json:"Labels"`
		} `json:"config"`
		Architecture string `json:"architecture"`
		Os           string `json:"os"`
		Rootfs       struct {
			Type    string   `json:"type"`
			DiffIds []string `json:"diff_ids"`
		} `json:"rootfs"`
		History []struct {
			Created    time.Time `json:"created"`
			CreatedBy  string    `json:"created_by"`
			EmptyLayer bool      `json:"empty_layer,omitempty"`
		} `json:"history"`
	} `json:"Docker"`
	DefaultMountsFilePath string `json:"DefaultMountsFilePath"`
	Isolation             string `json:"Isolation"`
	NamespaceOptions      []struct {
		Name string `json:"Name"`
		Host bool   `json:"Host"`
		Path string `json:"Path"`
	} `json:"NamespaceOptions"`
	Capabilities     interface{} `json:"Capabilities"`
	ConfigureNetwork string      `json:"ConfigureNetwork"`
	CNIPluginPath    string      `json:"CNIPluginPath"`
	CNIConfigDir     string      `json:"CNIConfigDir"`
	IDMappingOptions struct {
		HostUIDMapping bool          `json:"HostUIDMapping"`
		HostGIDMapping bool          `json:"HostGIDMapping"`
		UIDMap         []interface{} `json:"UIDMap"`
		GIDMap         []interface{} `json:"GIDMap"`
	} `json:"IDMappingOptions"`
	History []struct {
		Created    time.Time `json:"created"`
		CreatedBy  string    `json:"created_by"`
		EmptyLayer bool      `json:"empty_layer,omitempty"`
	} `json:"History"`
	Devices interface{} `json:"Devices"`
}
