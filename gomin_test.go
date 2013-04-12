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

func Test_Minimizejs(t *testing.T) {
	content, err := ioutil.ReadFile("exemple/js.js")
	if err != nil {
		t.Fatal(err)
	}
	min_c := MinJS(content)
	ioutil.WriteFile("exemple/js-min.js", min_c, 0777)
}
