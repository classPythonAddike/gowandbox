package gowandbox

import (
	"strings"
	"testing"
)

func TestGetLanguages(t *testing.T) {
	result, err := GetLanguages(10000)

	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("Got %v languages", len(result))
	}
}

func TestGetLanguagesTimeoutError(t *testing.T) {
	_, err := GetLanguages(1)

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}

	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatal(err.Error())
	}
}
