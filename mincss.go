package gomin

import (
	"strings"
)

func MinCSS(s []byte) []byte {
	return []byte(strings.Replace(strings.Replace(string(s), " ", "", -1), "\n", "", -1))
}
