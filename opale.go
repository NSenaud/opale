package opale

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type Mem struct {
	gorm.Model
	UsedPercent         float64
	Total, Used, Cached uint64
}

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

func GetMem() *Mem {
	v, _ := mem.VirtualMemory()

	mem := Mem{
		UsedPercent: v.UsedPercent,
		Total:       v.Total,
		Used:        v.Used,
		Cached:      v.Cached,
	}

	return &mem
}

func GetCpu() (*Cpu, *[]LogicalCore) {
	c, err := cpu.Percent(time.Second, true)
	checkErr(err)
	// FIXME Does not return usage per core
	p, err := cpu.Percent(time.Second, false)
	checkErr(err)
	f, err := cpu.Info()
	checkErr(err)

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

func InsertIntoDb(cpu *Cpu, threads *[]LogicalCore, mem *Mem) {
	db, err := gorm.Open("sqlite3", "./opale.db")
	checkErr(err)
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Cpu{})
	db.AutoMigrate(&LogicalCore{})
	db.AutoMigrate(&Mem{})

	// Create
	db.Create(cpu)
	for _, thread := range *threads {
		db.Create(&thread)
	}
	db.Create(mem)

	// Read
	var c Cpu
	var m Mem
	db.Last(&c)
	db.Last(&m)
	log.Printf("Last values inserted:\n\tCPU: %.02f%s\n\tRAM:%.02f%s",
		c.UsedPercent, "%", m.UsedPercent, "%")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
