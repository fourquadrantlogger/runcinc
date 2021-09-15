package rflag

import (
	"regexp"
	"strings"
)

const (
	LongflagPattern  = `^--[A-Za-z0-9][A-Za-z0-9_\.\-]*$`
	ShortflagPattern = `^-[A-Za-z0-9][A-Za-z0-9_\.\-]*$`
)

var (
	ShortKVflagPattern = ShortflagPattern[:len(ShortflagPattern)-1] + `=\S+`
	LongKVflagPattern  = LongflagPattern[:len(LongflagPattern)-1] + `=\S+`
)

type IndexFlag func(args []string) (idx int)

func stringshas(strs []string, s string) (idx int) {
	for i := 0; i < len(strs); i++ {
		if strs[i] == s {
			return i
		}
	}
	return -1
}

func matchFlagType(arg string) (t int) {
	kmatch, _ := regexp.MatchString(LongflagPattern, arg)
	if kmatch {
		return 1
	}
	kmatchshort, _ := regexp.MatchString(ShortflagPattern, arg)
	if kmatchshort {
		return 2
	}
	kvmatch, _ := regexp.MatchString(LongKVflagPattern, arg)
	if kvmatch {
		return 3
	}
	kvmatchshort, _ := regexp.MatchString(ShortKVflagPattern, arg)
	if kvmatchshort {
		return 4
	}
	return
}

//ParseFlag
//如果flag存在重复的key，appendkey中的key对value进行append，其它key的情况则后续的value直接丢弃而是选择第一个value
func ParseFlag(args []string, appendKey []string) (flags map[string][]string, unknownArgs []string) {
	flagargs := args
	flags = make(map[string][]string)
	for i := 0; i < len(flagargs); i++ {

		var k string
		var v string

		switch matchFlagType(flagargs[i]) {
		case 0:
			unknownArgs = append(unknownArgs, flagargs[i])
			continue
		case 1:
			k = flagargs[i][2:]
		case 2:
			k = flagargs[i][1:]
		case 3:
			kv := flagargs[i][2:]
			sep := strings.Index(kv, "=")
			k = kv[:sep]
			v = kv[sep+1:]
			if stringshas(appendKey, k) >= 0 {
				flags[k] = append(flags[k], v)
			} else {
				if len(flags[k]) == 0 {
					flags[k] = []string{v}
				}

			}
			continue
		case 4:
			kv := flagargs[i][1:]
			sep := strings.Index(kv, "=")
			k = kv[:sep]
			v = kv[sep+1:]
			if stringshas(appendKey, k) >= 0 {
				flags[k] = append(flags[k], v)
			} else {
				if len(flags[k]) == 0 {
					flags[k] = []string{v}
				}
			}
			continue
		}

		if i != len(flagargs)-1 {
			nextFlagType := matchFlagType(flagargs[i+1])
			if nextFlagType > 0 {
				flags[k] = []string{}
				//existflag
			} else {
				v = flagargs[i+1]
				i++
				if stringshas(appendKey, k) >= 0 {
					flags[k] = append(flags[k], v)
				} else {
					if len(flags[k]) == 0 {
						flags[k] = []string{v}
					}
				}
			}
		} else {
			flags[k] = []string{}
		}
	}
	return
}
