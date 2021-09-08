package main

import (
	"flag"
	"log"
	"os"

	"github.com/r0x0d/testgen/cmd/build"
	"github.com/r0x0d/testgen/pkg/templates"
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
	err := build.Prepare(flags)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	cases := build.Build()
	err = templates.BuildTemplate(cases)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
}
