package main

import (
	"fmt"
)

func PrinterError(str string) string {
	stringLen := len(str)
	numOfErrorChar := 0

	for _, char := range str {
		if char > rune('m') || char < rune('a') {
			numOfErrorChar++
		}
	}
	result := fmt.Sprintf("%d/%d", numOfErrorChar, stringLen)
	return  result
}

func main() {
	s1 :="aaaxbbbbyyhwawiwjjjwwm"
	fmt.Println(PrinterError(s1))

}