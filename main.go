package main

import (
	"fmt"

	"github.com/t-289/archive/hsfile"
)

func main() {
	ghs := hsfile.HashFile("README.md")
	ghs2 := hsfile.HashFile("/home/tiago/README.md")

	fmt.Printf("Hash: %x \n", ghs.Sum(nil))
	fmt.Printf("Hash: %x \n", ghs2.Sum(nil))
}
