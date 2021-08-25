package gowandbox

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	templ, err := GetTemplate("cpython", 10000)

	if err != nil {
		t.Error(err)
	}

	t.Logf("Got template for cpython - %v", templ)
}
