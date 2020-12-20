package gogethelp_test

import (
	"os"
	"testing"

	h "github.com/apichef/gogethelp"
)

func TestEnv(t *testing.T) {
	expected := "testing"
	os.Setenv("APP_ENV", expected)

	if got := h.Env("APP_ENV", "local"); got != expected {
		t.Errorf("Expected Env to return %s, revieved %s", expected, got)
	}
}
