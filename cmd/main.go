package main

import (
	"fmt"
	"io"

	"github.com/chzyer/readline"
)

func main() {
	input, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m#\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		InterruptPrompt: "^C",
	})
	if err != nil {
		panic(err)
	}

	defer input.Close()

	fmt.Println("This is Goldfoyle.")

prompt:
	for {
		line, err := input.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Error while reading line: %s", err)
			continue prompt
		}
	}
}
