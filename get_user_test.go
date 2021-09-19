package gowandbox

import (
	"context"
	"testing"
	"time"
  "strings"
)

func TestUserNotFoundError(t *testing.T) {
	user, err := GetUser("", context.Background()) // Dummy string
	if err != nil {
		t.Error(err.Error())
	}

	if user.Username != "" {
		t.Errorf("Got username - %v", user.Username)
	}

	t.Log("User was not found, as expected")
}

func TestUserTimeoutError(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)

	defer cancel()

	_, err := GetUser("", ctx)

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}

	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Error(err.Error())
	}

	t.Log("Request timed out as expected")
}
