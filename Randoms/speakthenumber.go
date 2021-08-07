package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//your code goes here

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type three numbers: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	num := input

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println(input)
	}
}
