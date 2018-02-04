package sensors

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/shirou/gopsutil/mem"
)

type Ram struct {
	gorm.Model
	UsedPercent         float64
	Total, Used, Cached uint64
}

func GetRam() *Ram {
	v, _ := mem.VirtualMemory()

	ram := Ram{
		UsedPercent: v.UsedPercent,
		Total:       v.Total,
		Used:        v.Used,
		Cached:      v.Cached,
	}

	return &ram
}
