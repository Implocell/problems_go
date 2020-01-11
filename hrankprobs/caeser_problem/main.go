package main

import (
	"fmt"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	// alphabethLower := "abcdefghijklmnopqrstuvwxyz"
	// alphabethUpper := strings.ToUpper(alphabethLower)

	var ret []rune
	for _, ch := range input {
		ret = append(ret, chiper(ch, delta))
		// switch {
		// case strings.IndexRune(alphabethLower, ch) >= 0:
		// 	ret = ret + string(rotate(ch, delta, []rune(alphabethLower)))
		// case strings.IndexRune(alphabethUpper, ch) >= 0:
		// 	ret = ret + string(rotate(ch, delta, []rune(alphabethUpper)))
		// default:
		// 	ret = ret + string(ch)
		// }
	}
	fmt.Println(string(ret))
}
func chiper(r rune, delta int) rune {
	if r > 'A' && r <= 'Z' {
		return rotateWithBase(r, 'A', delta)
	}
	if r > 'a' && r <= 'z' {
		return rotateWithBase(r, 'a', delta)
	}
	return r
}

func rotateWithBase(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)

}

// func rotate(s rune, delta int, key []rune) rune {
// 	idx := strings.IndexRune(string(key), s)
// 	// solving more manually
// 	// idx := -1
// 	// for i, r := range key {
// 	// 	if r == s {
// 	// 		idx = i
// 	// 		break
// 	// 	}
// 	// }
// 	if idx < 0 {
// 		panic("idx < 0")
// 	}

// 	idx = (idx + delta) % len(key)
// 	// looping version
// 	// for i := 0; i < delta; i++ {
// 	// 	idx++
// 	// 	if idx >= len(key) {
// 	// 		idx = 0
// 	// 	}
// 	// }

// 	return key[idx]
// }
