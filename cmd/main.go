package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

func main() {
	l, err := readline.NewEx(&readline.Config{})
	if err != nil {
		panic(err)
	}

	defer l.Close()

	fmt.Println("This is Goldfoyle.")
}
