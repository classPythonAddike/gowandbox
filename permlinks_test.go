package gowandbox

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestGetPermLink(t *testing.T) {

	result, err := GetPermLink("ia8loUVGVV8widMw", context.Background())

	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("ISO-8601 time - %v", result.Parameter.CreatedAt)
}

func TestGetPermLinkTimeoutError(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	_, err := GetPermLink("ia8loUVGVV8widMw", ctx)

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}

	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatal(err.Error())
	}

	t.Log("Request timed out, as expected")
}

func TestGetPermLinkBadLinkError(t *testing.T) {
	_, err := GetPermLink("abc", context.Background())

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}

	if !strings.Contains(err.Error(), "500 Error") {
		t.Fatal(err.Error())
	}

	t.Log("Got 500 error, as expected")
}
