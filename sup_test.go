package sup

import (
	"regexp"
	"testing"
)

func TestVersionUsesSemver(t *testing.T) {
	m, err := regexp.Match("^\\d+\\.\\d+\\.\\d+$", []byte(Version))
	if err != nil {
		t.Error("should not have returned error")
	}
	if !m {
		t.Error("version needs to be in semver")
	}
}
