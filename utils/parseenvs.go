package utils

import "strings"

func ParseEnvs(envs []string) (envmap map[string]string) {
	envmap = make(map[string]string)
	for _, kv_ := range envs {
		splitIndex := strings.IndexAny(kv_, "=")
		if splitIndex > 0 {
			envmap[kv_[:splitIndex]] = kv_[splitIndex+1:]
		}
	}
	return
}
