package containerimage

type Image struct {
	Env   []string `json:"Env"`
	Cmd   []string `json:"Cmd"`
	Lower []string
	Image string `json:"Image"`
}

type ImageDriver interface {
	Spec(image string) *Image
	Pull(image string)
}

var defaultImageDriver ImageDriver

func Driver() ImageDriver {
	return defaultImageDriver
}
func SetDriver(driver ImageDriver) {
	defaultImageDriver = driver
}
