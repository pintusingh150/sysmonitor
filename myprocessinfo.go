package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func GetRunningProcesses() string {
	var processes []*process.Process
	var err error
	processes, err = process.Processes()
	if err != nil {
		fmt.Println("Error in getting running processes", err)
		return ""
	}
	var processInfo string
	var processCounter int

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			name = "N/A"
		}
		pid := p.Pid
		processInfo += fmt.Sprintf("\nProcess Name: %s, PID: %d\n", name, pid)
		processCounter++

		if processCounter == 2 {
			break
		}
	}
	return processInfo
}
