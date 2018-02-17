package ram_test

import (
	"testing"

	"github.com/NSenaud/opale"
)

func TestGetRam(t *testing.T) {
	ram := opale.sensors.GetRam()

	if ram.UsedPercent < 0 || ram.UsedPercent > 100 {
		t.Error("Unexpected used memory percentage")
	}
}
