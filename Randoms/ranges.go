package main

import "fmt"

func main() {
	var b []int = []int{5, 10, 55, 6, 12}

	for i, element := range b {
		fmt.Printf("%d:%d\n", i, element)

	}
	var bb []int = []int{59, 16, 52, 60, 2}

	for z, element := range bb {
		fmt.Printf("%d:%d\n", z, element)
	}

}
