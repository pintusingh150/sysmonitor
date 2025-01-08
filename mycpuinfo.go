package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCpuUsage() float64 {
	var cpuUsage []float64
	var err error
	cpuUsage, err = cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error in getting CPU usage", err)
		return -1
	}
	if len(cpuUsage) > 0 {
		return cpuUsage[0]
	}
	return -1
}
