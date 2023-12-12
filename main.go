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

	day3 := newD3("input/3.txt")
	fmt.Printf("d3p1: %d\n", day3.part1())
	fmt.Printf("d3p2: %d\n", day3.part2())

	day4 := newD4("input/4.txt")
	fmt.Printf("d4p1: %d\n", day4.part1())
	fmt.Printf("d4p2: %d\n", day4.part2())

	day5 := newD5("input/5.txt")
	fmt.Printf("d5p1: %d\n", day5.part1())
	//fmt.Printf("d5p2: %d\n", day5.part2())
	fmt.Printf("d5p2: omitted\n")

	day6 := newD6("input/6.txt")
	fmt.Printf("d6p1: %d\n", day6.part1())
	fmt.Printf("d6p2: %d\n", day6.part2())

	day7 := newD7("input/7.txt")
	fmt.Printf("d7p1: %d\n", day7.part1())
	fmt.Printf("d7p2: %d\n", day7.part2())

	day8 := newD8("input/8.txt")
	fmt.Printf("d8p1: %d\n", day8.part1())
	fmt.Printf("d8p2: %d\n", day8.part2())

	day9 := newD9("input/9.txt")
	fmt.Printf("d9p1: %d\n", day9.part1())
	fmt.Printf("d9p2: %d\n", day9.part2())
}
