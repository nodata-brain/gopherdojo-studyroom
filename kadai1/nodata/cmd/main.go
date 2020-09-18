package main

import (
	"flag"

	"kadai1/pkg/filer"
)

func main() {
	dir := flag.String("dir", ".", "Search Dir")
	flag.Parse()
	f := filer.Filer{}
	f.SetPath(*dir)
	f.Filer()
}
