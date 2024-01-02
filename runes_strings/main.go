package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text := []rune(input)
	fmt.Println(CorrectString(&text))
}

func CorrectString(text *[]rune) string {
	fmt.Println(string(*text)[0])
	if unicode.IsLower((*text)[0]) {
		return "Wrong"
	}
	return "Right"
}
