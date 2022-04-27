package main

import (
	"fmt"
	"strconv"
)

func CountBits(number uint) int {
	var bit string
	var result int

	for ;; {
		if number % 2  == 0 {
			bit += strconv.Itoa(0)
		} else  {
			bit += strconv.Itoa(1)
		}
		if number == 1 {
			break
		}
		number = number / 2
	}
	//reverse but no need
	for i := len(bit); i > 0; i-- {
		fmt.Print(string(bit[i-1]))
		if string(bit[i-1]) == "1" {
			result++
		}
	}
	fmt.Println()
	return result
}

func main() {
	fmt.Println(CountBits(19))
}
