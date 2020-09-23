package main

import (
	"flag"

	"kadai1/pkg/img"
)

func main() {
	dir := flag.String("dir", ".", "Search Dir")
	flag.Parse()
	i := imgconv.Config{}
	i.SetPath(*dir)
	i.Conv()
}
