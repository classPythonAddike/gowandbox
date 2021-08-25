package gowandbox

import (
	"testing"
)

func TestUser(t *testing.T) {
	_, err := GetUser("", 10000) // Dummy string

	if err == nil {
		t.Error(err)
	}
}
