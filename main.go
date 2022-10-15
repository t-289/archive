package main

import (
	"fmt"
)

func main() {
	ghs := hashFile("README.md")
	ghs2 := hashFile("/home/tiago/README.md")

	fmt.Printf("Hash: %x \n", ghs.Sum(nil))
	fmt.Printf("Hash: %x \n", ghs2.Sum(nil))
}
