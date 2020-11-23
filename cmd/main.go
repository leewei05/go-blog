package main

import (
	"fmt"
	"os"
)

func main() {
	for true {
		fmt.Printf("go-sql > ")
		args := os.Args[1:]
		fmt.Println(args)

		if ok := checkExit(); ok {

		}
	}
}

func checkExit() bool {
	return false
}
