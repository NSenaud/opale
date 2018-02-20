package cpu

import (
	"time"

	"github.com/NSenaud/opale/db"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/shirou/gopsutil/cpu"
)

var log = logrus.New()

type Cpu struct{}

type CpuSnapshot struct {
	gorm.Model
	UsedPercent float64
	Microcode   string
	Cores       uint32
}

type CoreSnapshot struct {
	gorm.Model
	CoreId      uint32
	UsedPercent float64
	Mhz         float64
}

type ThreadSnapshot struct {
	gorm.Model
	ThreadId    uint32
	UsedPercent float64
	Mhz         float64
}

func (s *CpuSnapshot) Save() {
	db, err := gorm.Open("sqlite3", db.GetDbPath())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema if necessary
	db.AutoMigrate(&CpuSnapshot{})

	// Create
	db.Create(s)

	// TODO if debug
	// Read last input
	var c CpuSnapshot
	db.Last(&c)
	log.WithFields(logrus.Fields{
		"%":         c.UsedPercent,
		"Microcode": c.Microcode,
	}).Info("Inserted CPU values.")
}

//func (c *Cpu) Last() (cpu *CpuSnapshot) {
func Last() *CpuSnapshot {
	log.Debug("Opening connection with database...")
	db, err := gorm.Open("sqlite3", db.GetDbPath())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Debug("Requesting last cpu entry...")
	var cpu CpuSnapshot
	db.Last(&cpu)
	log.Debug("Done.")
	return &cpu
}

func New() (*CpuSnapshot, *[]ThreadSnapshot) {
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
		log.WithFields(logrus.Fields{
			"cpu_count":   threads,
			"info_length": len(f),
		}).Warn("len(f) != threads:")
		log.Panic("Threads count does not match []InfoStat array size.")
	}

	// Check core usage have all been returned.
	if len(p) != threads {
		log.Println("len(p) != threads:", len(p), "!=", threads)
		log.WithFields(logrus.Fields{
			"cpu_count":         threads,
			"percentage_length": len(p),
		}).Warn("Can't get used percentage per core.")
		// Avoid panicking due to last FIXME
		// log.Panic("Threads count does not match []float64 percentage array size.")
	}

	logicals := make([]ThreadSnapshot, threads)
	for i := 0; i < threads; i++ {
		logicals[i] = ThreadSnapshot{
			ThreadId: uint32(i),
			// FIXME Cf above
			// UsedPercent: p[i],
			UsedPercent: p[0],
			Mhz:         f[i].Mhz,
		}
	}

	cpu := CpuSnapshot{
		UsedPercent: c[0],
		Microcode:   f[0].Microcode,
		Cores:       uint32(threads),
	}

	return &cpu, &logicals
}
