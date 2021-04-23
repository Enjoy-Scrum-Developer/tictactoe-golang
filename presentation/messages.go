package presentation

import (
	"bufio"
	"fmt"
	"os"
)

const END = "\033[0m"
const RED = "\033[31m"
const GREEN = "\033[32m"
const YELLOW = "\033[33m"
const BLUE = "\033[34m"

func AlertMessage(message string) {
	fmt.Print(YELLOW, "\n", message, END, "\n")
}

func Prompt(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(GREEN, "\n", message, END)
	text, _ := reader.ReadString('\n')
	return text
}
