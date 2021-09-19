package gowandbox

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestGetLanguages(t *testing.T) {
	result, err := GetLanguages(context.Background())

	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("Got %v languages", len(result))
	}
}

func TestGetLanguagesTimeoutError(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)

	defer cancel()

	_, err := GetLanguages(ctx)

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}

	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatal(err.Error())
	}

	t.Log("Request timed out, as expected")
}
