package main

import (
	"fmt"
	"io"

	"github.com/chzyer/readline"
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m#\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		InterruptPrompt: "^C",
	})
	if err != nil {
		panic(err)
	}

	defer l.Close()

	fmt.Println("This is Goldfoyle.")

repl:
	for {
		line, err := l.Readline()
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
			fmt.Println("Error while reading line:", err)
			continue repl
		}
	}
}
