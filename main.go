package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1p1, d1p2 := d1()
	fmt.Printf("d1p1: %d\n", d1p1)
	fmt.Printf("d1p2: %d\n", d1p2)

	d2p1, d2p2 := d2()
	fmt.Printf("d2p1: %d\n", d2p1)
	fmt.Printf("d2p2: %d\n", d2p2)

	d3p1, d3p2 := d3()
	fmt.Printf("d3p1: %d\n", d3p1)
	fmt.Printf("d3p2: %d\n", d3p2)
}
