package gowandbox

import (
	"context"
	"strings"
	"testing"
)

func TestTemplate(t *testing.T) {
	template, err := GetTemplate("cpython", context.Background())

	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Got template for cpython - %v", template.Code)
}

func TestBadTemplateError(t *testing.T) {
	_, err := GetTemplate("abc", context.Background())

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}

	if !strings.Contains(err.Error(), "bad template_name") {
		t.Error(err.Error())
	}

	t.Log("Template for `abc` was not found, as expected")
}

// func TestTemplateTimeoutError(t *testing.T) {
// 	_, err := GetTemplate("", 1)
//
// 	if err == nil {
// 		t.Error("Got no error, but was expecting one!")
// 	}
//
// 	if !strings.Contains(err.Error(), "context deadline exceeded") {
// 		t.Error(err.Error())
// 	}
//
// 	t.Log("Request timed out as expected")
// }
