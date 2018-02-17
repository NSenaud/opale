package cpu_test

import (
	"testing"

	"github.com/NSenaud/opale"
)

func TestGetCpu(t *testing.T) {
	cpu, threads := opale.sensors.GetCpu()

	if cpu.UsedPercent < 0 || cpu.UsedPercent > 100 {
		t.Error("Unexpected cpu usage percentage")
	}

	for _, thread := range *threads {
		if thread.UsedPercent < 0 || thread.UsedPercent > 100 {
			t.Error("Unexpected thread usage percentage")
		}
	}
}
