package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Scan(&n)

	result := squareDigits(n)
	fmt.Println(result)
}

func squareDigits(numStr string) string {
	resultStr := ""

	for _, char := range numStr {
		digit := int(char - '0')
		squaredDigit := digit * digit
		resultStr += strconv.Itoa(squaredDigit)
	}

	return resultStr
}
