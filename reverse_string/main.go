package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Please type a word to test palindrom:")
	fmt.Scan(&s)

	if isPalindrom(s) {
		fmt.Println("This string is a palindrom:", s)
	} else {
		fmt.Println("This string is not a palindrom:", s)
	}

}

func reverseString(s string) string {
	var result string
	for _, letter := range s {
		result = string(letter) + result
	}
	return result
}

func isPalindrom(s string) bool {
	if strings.ToLower(s) == strings.ToLower(reverseString(s)) {
		return true
	} else {
		return false
	}
}
