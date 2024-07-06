package main

import (
	"fmt"
	"waveJVM/cmd"
)

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println("version 1.2.3")
	} else if command.HelpFlag || command.Class == "" {
		cmd.PrintUsage()
	} else {
		cmd.StartJVM(command)
	}
}
