package temp

import (
	"github.com/NSenaud/opale/db"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/shirou/gopsutil/host"
)

var log = logrus.New()

type Temp struct {
	SensorKey      string
	SensorKeyID    uint
	CelciusDegrees float64
}

type TemperatureSensor struct {
	gorm.Model
	Name      string
	Snapshots []TempSnapshot
}

type TempSnapshot struct {
	gorm.Model
	TemperatureSensorID uint
	CelciusDegrees      float64
}

func (s *Temp) Save() {
	db, err := gorm.Open("sqlite3", db.GetDbPath())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema if necessary
	db.AutoMigrate(&TemperatureSensor{})
	db.AutoMigrate(&TempSnapshot{})

	// Find existing sensor, if exists.
	var sensor TemperatureSensor

	// Else, create it.
	if db.Find(&sensor, "id = ?", s.SensorKeyID).RecordNotFound() {
		log.Debug("Create key in database")
		new_sensor := TemperatureSensor{Name: s.SensorKey}
		db.Create(&new_sensor)
	} else {
		log.WithFields(logrus.Fields{
			"name": sensor.Name,
		}).Debug("Existing sensor")
	}

	// Create
	db.Create(&TempSnapshot{
		TemperatureSensorID: s.SensorKeyID,
		CelciusDegrees:      s.CelciusDegrees,
	})

	// TODO if debug
	// Read last input
	var t TempSnapshot
	var r TemperatureSensor
	db.Last(&t).Related(&r)
	log.WithFields(logrus.Fields{
		"sensor": r.Name,
		"°C":     t.CelciusDegrees,
	}).Info("Inserted temp values.")
}

func Last(sensorName *string) *Temp {
	log.Debug("Opening connection with database...")
	db, err := gorm.Open("sqlite3", db.GetDbPath())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Debug("Requesting last temp entry for sensor %s", sensorName)
	var t TempSnapshot
	var r TemperatureSensor
	db.Last(&t).Related(&r)

	temp := Temp{
		SensorKey:      r.Name,
		SensorKeyID:    t.TemperatureSensorID,
		CelciusDegrees: t.CelciusDegrees,
	}

	return &temp
}

func New() *[]Temp {
	ts, err := host.SensorsTemperatures()
	if err != nil {
		panic(err)
	}

	temps := make([]Temp, 1)
	for i, t := range ts {
		temp := Temp{
			SensorKey:      t.SensorKey,
			SensorKeyID:    uint(i),
			CelciusDegrees: t.Temperature,
		}
		log.WithFields(logrus.Fields{
			"key":  temp.SensorKey,
			"temp": temp.CelciusDegrees,
		}).Debug("New sensor temperature.")
		temps = append(temps, temp)
	}

	return &temps
}
