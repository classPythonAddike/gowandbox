package gowandbox

import (
	"strings"
	"testing"
)

func TestUserNotFoundError(t *testing.T) {
	user, err := GetUser("", 10000) // Dummy string
	if err != nil {
		t.Error(err.Error())
	}

	if user.Username != "" {
		t.Errorf("Got username - %v", user.Username)
	}
}

func TestUserTimeoutError(t *testing.T) {
	_, err := GetUser("", 1)

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}

	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Error(err.Error())
	}
}
