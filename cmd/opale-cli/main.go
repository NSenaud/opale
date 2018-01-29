package main

import (
	"log"

	"github.com/NSenaud/opale"
)

func main() {
	cpu := opale.GetCpu()
	mem := opale.GetMem()

	log.Printf("CPU: %.02f%s", cpu.Combined, "%")
	i := 0
	for core := range cpu.PerCpu {
		log.Println("Core", i, ": ", core)
		i++
	}
	log.Printf("RAM: %.02f%s", mem.UsedPercent, "%")
}
