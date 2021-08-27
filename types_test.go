package gowandbox

import "testing"

func TestChangeURL(t *testing.T) {
	ChangeWandBoxUrl("abc")
	if WandBoxUrl != "abc" {
		t.Error("Could not change wandbox url used!")
	}
}

func TestResetURL(t *testing.T) {
	ResetWandBoxUrl()
	if WandBoxUrl != "https://wandbox.org/api/" {
		t.Error("Could not reset wandbox url used!")
	}
}
