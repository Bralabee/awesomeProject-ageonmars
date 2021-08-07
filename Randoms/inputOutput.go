package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	println("Line by Line")

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		tx := in.Text()
		fmt.Printf("txt = %s\n", in.Text())
		if strings.Trim(tx, "\n") == "next" {
			break
		}
	}
	//println("Raw Mode")
	//raw()
	println("Goodbye")
}

//func raw()  {
//	oldState, err := terminal.MakeRaw(0)
//	if err != nil{
//		log.Fatal(err.Error())
//	}
//	defer terminal.Restore(0, oldState)
//
//	reader := bufio.NewScanner(os.Stdin)
//	for{
//		c, _, err := reader.ReadRune()
//		if err != nil{
//			log.Fatal(err)
//		}
//		fmt.Printf("Rune = %c\n\r", c)
//		if c == 'q'{
//			return
//		}
//	}
//
//}
