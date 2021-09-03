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

//ParseFlag
//如果flag存在重复的key，appendkey中的key对value进行append，其它key的情况则value直接替换
func ParseFlag(args []string, appendKey []string) (flags map[string][]string, unknownArgs []string) {
	flagargs := args
	flags = make(map[string][]string)
	for i := 0; i < len(flagargs); i++ {
		kmatch, _ := regexp.MatchString(LongflagPattern, flagargs[i])
		kmatchshort, _ := regexp.MatchString(ShortflagPattern, flagargs[i])
		kvmatch, _ := regexp.MatchString(LongKVflagPattern, flagargs[i])
		kvmatchshort, _ := regexp.MatchString(ShortKVflagPattern, flagargs[i])
		var k string
		var v string
		if kmatch {
			k = flagargs[i][2:]
		} else if kmatchshort {
			k = flagargs[i][1:]
		} else if kvmatch {
			kv := flagargs[i][2:]
			sep := strings.Index(kv, "=")
			k = kv[:sep]
			v = kv[sep+1:]
			if stringshas(appendKey, k) >= 0 {
				flags[k] = append(flags[k], v)
			} else {
				flags[k] = []string{v}
			}
			continue
		} else if kvmatchshort {
			kv := flagargs[i][1:]
			sep := strings.Index(kv, "=")
			k = kv[:sep]
			v = kv[sep+1:]
			if stringshas(appendKey, k) >= 0 {
				flags[k] = append(flags[k], v)
			} else {
				flags[k] = []string{v}
			}
			continue
		} else {
			unknownArgs = append(unknownArgs, flagargs[i])
			continue
		}

		if i != len(flagargs)-1 {
			vmatchshort, _ := regexp.MatchString(ShortflagPattern, flagargs[i+1])
			vmatch, _ := regexp.MatchString(LongflagPattern, flagargs[i+1])
			if vmatchshort || vmatch {
				flags[k] = []string{}
				//existflag
			} else {
				v = flagargs[i+1]
				i++
				if stringshas(appendKey, k) >= 0 {
					flags[k] = append(flags[k], v)
				} else {
					flags[k] = []string{v}
				}
			}
		} else {
			flags[k] = []string{}
		}
	}
	return
}
