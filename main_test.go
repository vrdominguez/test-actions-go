package main

import (
	"os"
	"testing"
)

func TestGetCurrentOS(t *testing.T) {
	envOS, ok := os.LookupEnv("os")

	if !ok {
		envOS = "linux"
	}

	if currentOS := getOS(); currentOS != envOS {
		t.Errorf("FAIL: Expected system: %s, detected: %s", envOS, currentOS)
	} else {
		t.Logf("OK: System detected as expected: %s", envOS)
	}

	// Now test a incorrect OS
	envOS = "VictorOS"

	if currentOS := getOS(); currentOS != "VictorOS" {
		t.Logf("OK: Expected system: %s, detected: %s", envOS, currentOS)
	} else {
		t.Errorf("FAIL: System detected as %s", envOS)
	}
}
