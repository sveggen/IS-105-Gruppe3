package main

import (
	"os"
	"./lineshift"
)

func main() {
	filename := os.Args[1]
	lineshift.SearchForLineshift(filename)
}