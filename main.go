package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	// => returns "WeIrD StRiNg CaSe"
	var result = ""

	//字符串str开头没有空格，即假设其初始值为-1
	indexOfSpace := -1
	wordIndex := 0
	for index, char := range str {
		wordIndex =  index - (indexOfSpace + 1)

		if wordIndex % 2 == 0 {
			result += strings.ToUpper(string(char))
		} else {
			result += strings.ToLower(string(char))
		}
		if char == ' ' {
			indexOfSpace = index
		}
	}
	return result
}


func toWeirdCase2(str string) string {
	var result = ""

	//str1 := "Weird string case"
	strArray := strings.Split(str, string(' '))
	for _, word := range strArray {
		for wordIndex, char := range word {
			if wordIndex % 2 == 0 {
				result += strings.ToUpper(string(char))
			} else {
				result += strings.ToLower(string(char))
			}

		}
		if len(strArray) > 1 {
			result += string(' ')
		}
	}
	return result
	
}



func main() {
	str1 := "Weird string casE HellO"
	str2 := "WeirDD"
	fmt.Println(toWeirdCase(str1))
	fmt.Println(toWeirdCase2(str1))
	fmt.Println(toWeirdCase(str2))
	fmt.Println(toWeirdCase2(str2))
}
