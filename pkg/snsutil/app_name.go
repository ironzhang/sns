package snsutil

import "strings"

func ParseAppName(name string) (cluster, service string) {
	ss := strings.SplitN(name, ".", 2)
	if len(ss) < 2 {
		return "default", name
	}
	return ss[0], ss[1]
}
