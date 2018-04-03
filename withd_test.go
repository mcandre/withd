package withd_test

import (
	"testing"

	"github.com/mcandre/withd"
)

func TestVersion(t *testing.T) {
	if withd.Version == "" {
		t.Errorf("Expected withd version to be non-blank")
	}
}
