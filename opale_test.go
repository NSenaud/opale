package opale_test

import (
	"testing"

	"github.com/NSenaud/opale"
)

func TestGetMem(t *testing.T) {
	mem := opale.GetMem()

	if mem.UsedPercent < 0 || mem.UsedPercent > 100 {
		t.Error("Unexpected used memory percentage")
	}
}
