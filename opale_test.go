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

func TestGetCpu(t *testing.T) {
	cpu, threads := opale.GetCpu()

	if cpu.UsedPercent < 0 || cpu.UsedPercent > 100 {
		t.Error("Unexpected cpu usage percentage")
	}

	for _, thread := range *threads {
		if thread.UsedPercent < 0 || thread.UsedPercent > 100 {
			t.Error("Unexpected thread usage percentage")
		}
	}
}
