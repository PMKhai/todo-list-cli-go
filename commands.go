package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var todoList = make(map[int]string)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to clear screen: %s", err)
	}
}

func addCmd(args []string) {
	if checkArgs(args, 1) {
		todoList[len(todoList)] = args[1]
	}
}

func editCmd(args []string) {
	if checkArgs(args, 2) {
		replace(args[1], args[2])
	}
}

func rmCmd(args []string) {
	if checkArgs(args, 1) {
		remove(args[1])
	}
}

func lsCmd(args []string) {
	items := make([]string, 0, len(todoList))
	for _, item := range todoList {
		items = append(items, item)
	}
	fmt.Println(strings.Join(items, ", "))
}

func clearCmd(args []string) {
	clearScreen()
}

func replace(oldItem, newItem string) {
	for i, item := range todoList {
		if item == oldItem {
			todoList[i] = newItem
			break
		}
	}
}

func remove(rmItem string) {
	for i, item := range todoList {
		if item == rmItem {
			delete(todoList, i)
			break
		}
	}
}

func checkArgs(args []string, required int) bool {
	if len(args) < required+1 || args[required] == "" {
		fmt.Printf("Error: %s is missing some parameter", args[0])
		return false
	}
	return true
}

func executeCLI(args []string) {
	commandMap := map[string]func([]string){
		"add":   addCmd,
		"edit":  editCmd,
		"rm":    rmCmd,
		"ls":    lsCmd,
		"clear": clearCmd,
	}

	if fn, ok := commandMap[args[0]]; ok {
		fn(args)
	} else {
		fmt.Println("Error: " + args[0])
	}
}
