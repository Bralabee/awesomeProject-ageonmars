package main

import "fmt"

func main() {
	var a []int = []int{5, 6, 8, 9, 10}
	fmt.Println(cap(a))
}
