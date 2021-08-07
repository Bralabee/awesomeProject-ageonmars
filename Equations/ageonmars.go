package main

import "fmt"

func calcAge(earthAge int) int {
	earthdays := earthAge * 365
	marsAge := earthdays / 687

	return marsAge
}

func main() {
	earthAge := 20

	mars := calcAge(earthAge)
	fmt.Println(mars)
}
