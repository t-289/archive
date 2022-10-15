package main

import (
	"fmt"
)

func main() {
	ghs := hsfile.hashFile("README.md")
	ghs2 := hsfile.hashFile("/home/tiago/README.md")

	fmt.Printf("Hash: %x \n", ghs.Sum(nil))
	fmt.Printf("Hash: %x \n", ghs2.Sum(nil))
}
