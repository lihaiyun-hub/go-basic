package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(SingleNumber(nums))
}

func SingleNumber(nums []int) int {
	counter := make(map[int]int)
	for _, value := range nums {
		counter[value]++

	}
	for key, value := range counter {
		if value == 1 {
			return key
		}
	}
	return 0
}
