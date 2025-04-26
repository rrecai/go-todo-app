package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("To-Do App Starts!")
	fmt.Println("Enter 'Help' to see commands.")
	var x []string
	var y string
	var command string
	file, _ := os.Open("tasks.csv")
	defer file.Close()

	filescan := bufio.NewScanner(file)
	for filescan.Scan() {
		x = []string(strings.Split(filescan.Text(), ", "))
	}

	for {
		fmt.Println("Enter Command:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		y = scanner.Text()
		command = strings.Split(y, " ")[0]
		switch command {
		case "Help":
			fmt.Println("'Add TASK', 'Delete TASK', 'List', 'Write' and 'Exit'")
		case "List":
			fmt.Println(strings.Join(x, ", "))
		case "Exit":
			os.Exit(0)
		case "Add":
			x = append(x, strings.TrimPrefix(y, "Add "))
		case "Delete":
			x = slices.DeleteFunc(x, func(z string) bool {
				return z == strings.TrimPrefix(y, "Delete ")
			})
		case "Write":
			os.WriteFile("tasks.csv", []byte(strings.Join(x, ", ")), 0644)
		}
	}
}
