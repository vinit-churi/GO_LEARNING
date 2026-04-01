//go:build solution
// +build solution

package main

import "fmt"

func ByteLength(text string) int {
	return len(text)
}

func FirstRune(text string) rune {
	runes := []rune(text)
	return runes[0]
}

func ReverseRunes(text string) string {
	runes := []rune(text)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

func main() {
	fmt.Println(ByteLength("Go语言"))
	fmt.Println(string(FirstRune("你好")))
	fmt.Println(ReverseRunes("Go语言"))
}
