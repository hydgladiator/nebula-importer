package main

import (
	"flag"
	"log"

	"github.com/yixinglu/nebula-importer/pkg/cmd"
)

var configuration = flag.String("config", "", "Specify importer configure file path")

func main() {
	flag.Parse()

	if err := cmd.Run(*configuration); err != nil {
		log.Fatal(err)
	}
}
