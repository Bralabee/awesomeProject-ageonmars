package main

import (
	"fmt"
	_ "github.com/gorilla/mux"
)

//Processor struct
type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoryCapacity int
	memoryType     string
}

type Computer struct {
	Processor
	Memory
	price int
}

func main() {

	computer := Computer{}
	computer.price = 50000
	computer.processorName = "intel i7"
	computer.cores = 6
	computer.memoryType = "DDR4"

	computer1 := Computer{
		Processor: Processor{
			processorName: "intel i12",
			cores:         20,
		},
		Memory: Memory{
			memoryCapacity: 60,
			memoryType:     "plenty",
		},
		price: 10000,
	}

	fmt.Println(computer)
	fmt.Println(computer1)
	writeMessage("Hello People")

}
func writeMessage(msg string) {
	fmt.Println(msg)
}
