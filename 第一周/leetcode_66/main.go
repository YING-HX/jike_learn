package main

import "fmt"

// 加1
func main() {
	a := plusOne([]int{1, 4, 6})
	fmt.Printf("a= %v", a)
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		num := digits[i] + 1
		if num == 10 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			return digits
		}
	}
	return digits
}
