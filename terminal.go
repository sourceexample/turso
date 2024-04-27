package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testturso/modTurso"
)

func inputhandler() {
	reader := bufio.NewReader(os.Stdin)
	bQuit := false
	for !bQuit {
		fmt.Print(">: ")
		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)
		text = strings.Trim(text, "\r")
		text = strings.Trim(text, "\n")
		command := strings.Split(text, " ")
		commandlength := len(command)
		if commandlength < 1 {
			continue
		}

		switch command[0] {
		case "quit":
			bQuit = true
			break
		case "exit":
			bQuit = true
			break
		case "add":
			if commandlength < 2 {
				fmt.Println("add project error, params not enough")
				continue
			}
			err := modTurso.GetSingleProject().Insert(command[1])
			if err != nil {
				fmt.Println(err)
			}
		case "query":
			var name string = ""
			if commandlength > 1 {
				name = command[1]
			}
			err := modTurso.GetSingleProject().Query(name)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
