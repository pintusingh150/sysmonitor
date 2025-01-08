package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the entry point of the program")

	cpuUsage := GetCpuUsage()
	memUsage := GetMemUsage()
	pInfo := GetRunningProcesses()

	fmt.Println("CPU Usage: ", cpuUsage)
	fmt.Println("Memory Usage: ", memUsage)
	fmt.Println("Running Processes: ", pInfo)
}
