package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	// => returns "WeIrD StRiNg CaSe"
	// => returns "WeIrD StRiNg CaSe"
	// => returns "WeIrD StRiNg CaSe"
	// => returns "WeIrD StRiNg CaSe"
	var result = ""

	indexOfSpace := -1
	wordIndex := 0
	for index, char := range str {
		wordIndex =  index - (indexOfSpace + 10)

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







func main() {
	str1 := "Weird string case"
	fmt.Println(toWeirdCase(str1))

}
