package cmd

import (
	"flag"
	"fmt"
	"machine/internal/machine"
	"os"
)

func Run() {
	// Load arguments from the user
	var objectFile string
	var helpFlag bool
	var cpuFreq int
	flag.StringVar(&objectFile, "load", "", "binary file to load")
	flag.BoolVar(&helpFlag, "help", false, "show usage and exit")
	flag.IntVar(&cpuFreq, "clockspeed", 8, "(optional) clock speed in Mhz")
	flag.Parse()

	if objectFile == "" || helpFlag {
		flag.Usage()
		os.Exit(1)
	}

	// initialise machine
	m := machine.Boot()

	// load file
	err := m.LoadProgram(objectFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Run the cpu
	err = m.RunCpu()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

}
