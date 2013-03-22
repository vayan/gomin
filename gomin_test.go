package gomin

import (
	// "bitbucket.org/vayan/gomin"
	"io/ioutil"
	"testing"
)

func Test_Minimizecss(t *testing.T) {
	content, err := ioutil.ReadFile("exemple/bootstrap.css")
	if err != nil {
		t.Fatal(err)
	}
	min_c := MinCSS(content)
	ioutil.WriteFile("exemple/bootstrap-min.css", min_c, 0777)
}
