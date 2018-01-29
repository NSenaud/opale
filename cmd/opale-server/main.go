package main

import (
	"github.com/NSenaud/opale"
)

func main() {
	for {
		cpu, threads := opale.GetCpu()
		mem := opale.GetMem()

		opale.InsertIntoDb(cpu, threads, mem)
		// Sleep is not necessary yet since we already wait for a second in
		// GetCpu(), however it will be necessary as soon as we will be async,
		// and the interval will be setup in configuration file.
	}
}
