package cic

type CicConfig struct {
	Env             []string `json:"Env"`
	Cmd             []string `json:"Cmd"`
	ImagePullPolicy ImagePullPolicy
	Images          []string
	ImageRoot       string
	Name            string
	CicVolume       string
}
type ImagePullPolicy string

const (
	ImagePullPolicyfNotPresent ImagePullPolicy = "IfNotPresent"
	imagePullPolicyAlways      ImagePullPolicy = "Always"
)
