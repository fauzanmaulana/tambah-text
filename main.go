package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var run bool = true

func getInput(text string) string {
	fmt.Print(text + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	value := strings.Trim(input, "\n")
	return value
}

func inputText(input string, text func(string) string) string {
	output := text(input)
	return output
}

func printText(texts ...string) {
	for _, text := range texts {
		fmt.Println(text)
	}
}

func main() {
	// define slice
	var text = []string{"fauzan maulana"}
	var newText []string

	for run {
		input := inputText("please input some name", getInput)

		// append new slice
		newText = append(text, input)
		text = newText

		// define output
		printText(text...)

		confirm := inputText("want more? (y/n)", getInput)
		if confirm == "n" {
			run = false
		}
	}

	fmt.Println("program finish")
}
