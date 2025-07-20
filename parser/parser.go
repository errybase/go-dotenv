package parser

import (
	"regexp"
)

var re = regexp.MustCompile(`(?m)^(?P<key>.+?)\s*=\s*(("(?P<value1>.+)")|(?P<value>.+))$`)

func Parse(b []byte) map[string]string {
	h := make(map[string]string)
	for _, m := range re.FindAllSubmatch(b, -1) {
		k := string(m[re.SubexpIndex("key")])
		v := string(m[re.SubexpIndex("value")])
		if v1 := string(m[re.SubexpIndex("value1")]); v1 != "" {
			v = v1
		}
		h[k] = v
	}
	return h
}
