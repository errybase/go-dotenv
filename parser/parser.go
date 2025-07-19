package parser

import (
	"regexp"
)

var pattern = regexp.MustCompile(`(?m)^(?P<key>.+?)\s*=\s*(?P<value>.+)$`)

func Parse(b []byte) map[string]string {
	v := make(map[string]string)
	for _, m := range pattern.FindAllSubmatch(b, -1) {
		v[string(m[1])] = string(m[2])
	}
	return v
}
