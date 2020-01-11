package main

import (
	"fmt"
	"strings"
)

func main() {
	input := CamelCase()
	answer := 1

	// solution 2
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}
	}

	// solution 1
	// for _, ch := range input {
	// 	min, max := 'A', 'Z'
	// 	if ch >= min && ch <= max {
	// 		answer++
	// 	}
	// }

	fmt.Println(answer)
}

func CamelCase() string {
	var input string
	fmt.Scanf("%s\n", &input)
	return input
}
