package utils

import (
	"sort"
	"strings"
)

func ParseEnvs(envs []string) (envmap map[string]string) {
	envmap = make(map[string]string)
	for _, kv := range envs {
		splitIndex := strings.IndexAny(kv, "=")
		if splitIndex > 0 {
			k := kv[:splitIndex]
			v := kv[splitIndex+1:]
			if _, h := envmap[k]; h {
				MergeEnv(envmap, []string{kv})
			} else {
				envmap[k] = v
			}

		}
	}
	return
}

func MergeEnv(envmap map[string]string, envs []string) {
	for _, kv := range envs {
		splitIndex := strings.IndexAny(kv, "=")
		if splitIndex > 0 {
			k := kv[:splitIndex]
			v := kv[splitIndex+1:]
			if oldv, have := envmap[k]; have {
				if strings.Contains(k, "PATH") {
					pathValues := strings.Split(oldv, ":")
					vValues := strings.Split(v, ":")
					pathValues = MergeStrings(pathValues, vValues)
					sort.Strings(pathValues)
					newV := strings.Join(pathValues, ":")
					v = newV
				}
				envmap[k] = v
				//overwrite
				//fmt.Println("overwrite",k,oldv,v)

			} else {
				envmap[k] = v
			}
		}
	}
}
