package main

import (
	"flag"
	"log"

	"github.com/gitbnw/kin/convert"
)

func main() {
	file := flag.String("file-path", "sample.txt", "path to file input: string")
	flag.Parse()

	if err := convert.Run(file); err != nil {
		log.Print(err)
	}
}
