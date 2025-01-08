package main

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func GetMemUsage() float64 {
	var memUsage *mem.VirtualMemoryStat
	var err error
	memUsage, err = mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error in getting memory usage", err)
		return -1
	}
	if memUsage != nil {
		return memUsage.UsedPercent
	}
	return -1
}
