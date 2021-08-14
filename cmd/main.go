package main

import (
	"flag"
	"github.com/r0x0d/testgen/cmd/build"
	"github.com/r0x0d/testgen/pkg/templates/python"
	"log"
	"os"
)

var flags = []flag.Flag{
	flag.Flag{
		Name:     "output",
		Usage:    "Sets the output test file path",
		DefValue: ".",
	},
}

func main() {
	log.Println(os.Args[1])
	err := build.Run(flags)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}

	cases := build.Build()
	err = python.BuildTemplate(cases)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	// ...
}
