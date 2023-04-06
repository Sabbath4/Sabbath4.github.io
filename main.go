package main
//Write a program that takes a string as input and prints the reverse of the string// 
import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "apple"
	reversedStr := reverseString(str)
	fmt.Println(reversedStr)
}