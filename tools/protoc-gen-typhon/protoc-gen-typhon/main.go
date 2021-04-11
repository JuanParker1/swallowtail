package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/monzo/wearedev/tools/protoc-gen-typhon/gen"
)

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println(gen.Version)
		os.Exit(0)
	}

	g := newGenerator()
	gen.Main(g)
}
