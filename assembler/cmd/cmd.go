package cmd

import (
	"assembler/pkg/compiler"
	"flag"
	"fmt"
	"os"
)

func Run() {
	var out string
	var in []string

	// register flags
	flag.StringVar(&out, "o", "", "object file")
	// parse
	flag.Parse()
	in = flag.Args()

	// validate flags
	if len(in) != 1 || out == "" {
		fmt.Fprintf(os.Stderr, "correct usage : %s -o <objfile> <sourcecode>\n", os.Args[0])
		os.Exit(1)
	}

	// run the compiler
	err := compiler.Assemble(in[0], out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
