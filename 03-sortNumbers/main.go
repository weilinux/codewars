package main

import "fmt"

func SortNumbers(numbers []int) []int {
	//bubbleSort排序吧
	//var numbers = []int{1, 20, 3, 10, 5}
	//round#1 1 3 10 5 20
	//round#2 1 3 5 10 20
	//round#N-1

	for j := len(numbers); j > 0; j-- {
		for i := 0; i< j -1; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}
	return numbers // your code here
}

func SortNumbers2(numbers []int) []int {
	//quickSort
	return []int{}
}


func main() {
	var numbers = []int{1, 20, 3, 10, 5, 50, 200, 1000, 2, 70, 250}
	fmt.Println(SortNumbers(numbers))
}
