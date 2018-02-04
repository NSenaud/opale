package db

import (
	"log"

	"github.com/NSenaud/opale/sensors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InsertIntoDb(cpu *sensors.Cpu, threads *[]sensors.LogicalCore, ram *sensors.Ram) {
	db, err := gorm.Open("sqlite3", "./opale.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&sensors.Cpu{})
	db.AutoMigrate(&sensors.LogicalCore{})
	db.AutoMigrate(&sensors.Ram{})

	// Create
	db.Create(cpu)
	for _, thread := range *threads {
		db.Create(&thread)
	}
	db.Create(ram)

	// Read
	var c sensors.Cpu
	var m sensors.Ram
	db.Last(&c)
	db.Last(&m)
	log.Printf("Last values inserted:\n\tCPU: %.02f%s\n\tRAM:%.02f%s",
		c.UsedPercent, "%", m.UsedPercent, "%")
}
