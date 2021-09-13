package cic

type CicConfig struct {
	CapAdd          []string
	CapDrop         []string
	Env             []string `json:"Env"`
	Volume          []string
	CopyParentEnv   bool
	Cmd             []string `json:"Cmd"`
	ImagePullPolicy ImagePullPolicy
	Images          []string
	ImageRoot       string
	Authfile        string
	Name            string
	CicVolume       string
}
type ImagePullPolicy string

const (
	ImagePullPolicyfNotPresent ImagePullPolicy = "IfNotPresent"
	imagePullPolicyAlways      ImagePullPolicy = "Always"
)
