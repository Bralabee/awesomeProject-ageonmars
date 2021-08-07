package main

import "fmt"

func main() {
	var a []int = []int{2, 2, 5, 6, 12, 8, 9, 7, 5, 4, 3, 88, 3, 9, 2}

	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}

}
