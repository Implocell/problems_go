package main

import (
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

func main() {
	var myNum string
	fmt.Println("Hello World")
	x := fmt.Sprint("hello world", 1, 2, 3)
	fmt.Scan(&myNum)
	fmt.Println(x)
	fmt.Println(myNum)
}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}

// func normalize(phone string) string {
// 	var buf bytes.Buffer
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}
// 	return buf.String()
// }

func returnNumber(number int) int {
	num := number
	return num * 2
}
