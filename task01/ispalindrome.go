package main

import (
	"fmt"
	"math"
)

func main() {
	

}

func Isalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	length := 0
	x_ := x
	for ; x_ > 0; x_ /= 10 {
		length++
	}
	fmt.Println(x)
	fmt.Println(length)
	for i := 0; i < length/2; i++ {

		high := (x / int(math.Pow10(length-1-i))) % 10
		low := (x / int(math.Pow10(i))) % 10
		if high != low {
			return false
		}

	}
	return true

}
