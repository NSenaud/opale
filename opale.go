package opale

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type Mem struct {
	UsedPercent         float64
	Total, Used, Cached uint64
}

func GetMem() *Mem {
	v, _ := mem.VirtualMemory()

	mem := Mem{UsedPercent: v.UsedPercent,
		Total:  v.Total,
		Used:   v.Used,
		Cached: v.Cached}

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	return &mem
}
