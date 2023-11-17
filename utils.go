package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getOsReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func loopFreeBirdTool(reader *bufio.Reader) {
	for {
		printToolPrompt()
		arg := readArgument(reader)
		if strings.TrimSpace(arg) == "exit" {
			break
		}
		executeCLI(strings.Split(strings.TrimSpace(arg), " "))
	}
}

func printToolPrompt() {
	fmt.Print("TodoList -> ")
}

func readArgument(reader *bufio.Reader) string {
	text, err := reader.ReadString('\n')
	handleError(err)
	return text
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
