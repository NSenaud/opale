package sensors

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/shirou/gopsutil/cpu"
)

type Cpu struct {
	gorm.Model
	UsedPercent float64
	Microcode   string
	Cores       uint32
}

type PhysicalCore struct {
	gorm.Model
	CoreId      uint32
	UsedPercent float64
	Mhz         float64
}

type LogicalCore struct {
	gorm.Model
	ThreadId    uint32
	UsedPercent float64
	Mhz         float64
}

func GetCpu() (*Cpu, *[]LogicalCore) {
	c, err := cpu.Percent(time.Second, true)
	if err != nil {
		panic(err)
	}
	// FIXME Does not return usage per core
	p, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	f, err := cpu.Info()
	if err != nil {
		panic(err)
	}

	threads, _ := cpu.Counts(true)
	log.Println("CPU have", threads, "threads.")

	// Check threads/cores infos have all been returned.
	if len(f) != threads {
		log.Println("len(f) != threads:", len(f), "!=", threads)
		log.Panic("Threads count does not match []InfoStat array size.")
	}

	// Check core usage have all been returned.
	if len(p) != threads {
		log.Println("len(p) != threads:", len(p), "!=", threads)
		// Avoid panicking due to last FIXME
		// log.Panic("Threads count does not match []float64 percentage array size.")
	}

	logicals := make([]LogicalCore, threads)
	for i := 0; i < threads; i++ {
		logicals[i] = LogicalCore{
			ThreadId: uint32(i),
			// FIXME Cf above
			// UsedPercent: p[i],
			UsedPercent: p[0],
			Mhz:         f[i].Mhz,
		}
	}

	cpu := Cpu{
		UsedPercent: c[0],
		Microcode:   f[0].Microcode,
		Cores:       uint32(threads),
	}

	return &cpu, &logicals
}
