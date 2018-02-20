package ram

import (
	"github.com/NSenaud/opale/db"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/shirou/gopsutil/mem"
)

var log = logrus.New()

type RamSnapshot struct {
	gorm.Model
	UsedPercent         float64
	Total, Used, Cached uint64
}

func (s *RamSnapshot) Save() {
	db, err := gorm.Open("sqlite3", db.GetDbPath())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema if necessary
	db.AutoMigrate(&RamSnapshot{})

	// Create
	db.Create(s)

	// TODO if debug
	// Read last input
	var r RamSnapshot
	db.Last(&r)
	log.WithFields(logrus.Fields{
		"%":      r.UsedPercent,
		"Total":  r.Total,
		"Used":   r.Used,
		"Cached": r.Cached,
	}).Info("Inserted RAM values.")
}

func Last() *RamSnapshot {
	db, err := gorm.Open("sqlite3", db.GetDbPath())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var ram RamSnapshot
	db.Last(&ram)
	return &ram
}

func New() *RamSnapshot {
	v, _ := mem.VirtualMemory()

	ram := RamSnapshot{
		UsedPercent: v.UsedPercent,
		Total:       v.Total,
		Used:        v.Used,
		Cached:      v.Cached,
	}

	return &ram
}
