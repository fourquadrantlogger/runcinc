package common

type Image struct {
	Env   []string `json:"Env"`
	Cmd   []string `json:"Cmd"`
	Lower []string
	Image string `json:"Image"`
}

type ImageDriver interface {
	Name() string
	Spec(image string) *Image
	Pull(image, registrySecret string) error
}
