package gowandbox

import (
	"testing"
)

func TestGetLanguages(t *testing.T) {
	result, err := GetLanguages(10000)

	if err != nil {
		t.Log(result)
		t.Error(err)
	} else {
		t.Logf("Got %v languages", len(result))
	}
}
