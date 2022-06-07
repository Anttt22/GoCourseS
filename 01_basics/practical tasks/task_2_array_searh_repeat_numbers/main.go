package main

import (
	"fmt"
)

func Search(array [10]int) string {
	var temp, counter int
	for i := 0; i < len(array); i++ {
		temp = array[i]
		for j := i + 1; j < len(array); j++ {
			if temp == array[j] {
				counter++
				if counter > 0 {
					return "true"
				}
			}
		}
	}
	return " no repeated numbers"
}

func main() {
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 4, 9, 10}
	nums1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("first array search results:", Search(nums))
	fmt.Println("second array search results:", Search(nums1))
	fmt.Println("end of programm")
}
