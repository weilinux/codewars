package main

import (
	"fmt"
	"strconv"
)

func CountBits(number uint) int {
	var bit string
	var result int

	for ; number > 0; number /= 2 {
		if (number % 2) != 0 {
			result++
			bit += strconv.Itoa(1)
		} else  {
			bit += strconv.Itoa(0)
		}
	}
	//reverse but no need
	for i := len(bit); i > 0; i-- {
		fmt.Print(string(bit[i-1]))
	}
	fmt.Println()
	return result
}

func main() {
	fmt.Println(CountBits(19))
}
