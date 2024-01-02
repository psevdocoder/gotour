package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b int

	_, err := fmt.Scan(&a)
	if err != nil {
		log.Println("Ошибка ввода первого числа:", err)
		return
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		log.Println("Ошибка ввода второго числа:", err)
		return
	}

	sum := a + b
	fmt.Println(sum)
}
