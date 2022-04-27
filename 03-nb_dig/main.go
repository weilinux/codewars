package  main

import (
	"fmt"
	"strconv"
)

func NbDig(n int, d int) int {
	var result = 0
	var numString = ""

	for i := 0; i<= n; i++ {
		numString += strconv.Itoa(i * i)
	}
	for _, char := range numString {
		if strconv.Itoa(d) == string(char) {
			result++
		}
	}
	return result
}



func main() {
	fmt.Println(NbDig(25, 1))
}
