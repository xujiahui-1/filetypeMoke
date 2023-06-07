package types

import "strings"

// splitMime 切割MIME的方法，返回前半段和后半段
func splitMime(s string) (string, string) {
	x := strings.Split(s, "/")
	if len(x) > 1 {
		return x[0], x[1]
	}
	//说明切不了。返回后半段为空
	return x[0], ""
}
