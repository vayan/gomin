package gomin

import (
	"strings"
)

func MinCSS(s []byte) []byte {
	return []byte(strings.Replace(string(s), "\n", " ", -1))
}
