package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kawakattsun/dotenveditor"
)

var (
	version  string
	revision string
)

func main() {
	var (
		out     string
		version bool
	)
	flag.BoolVar(&version, "v", false, "show dotenveditor version")
	flag.StringVar(&out, "o", ".env", "output filename")
	flag.Parse()
	if version {
		showVersion()
		os.Exit(0)
	}
	in := flag.Arg(0)
	if in == "" {
		fmt.Print("error: please specify the target dotenv file\n")
		os.Exit(1)
	}

	if err := dotenveditor.Run(in, out); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func showVersion() {
	fmt.Printf("dotenveditor version %s (%s)\n", version, revision)
}
