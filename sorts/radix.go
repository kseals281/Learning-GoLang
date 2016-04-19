package main

import "fmt"

func main() {
	nums := [10]int{23, 86, 3, 98, 25, 75, 100, 49, 33, 10}
	var nums2 [10]int
	iter := 0
	modnum := 10
	for i := 0; i < 30; i++ {
		current := 0
		for j := 0; j < 10; j++ {
			if (nums[j] % modnum) == iter {
				temp := nums[current]
				nums[current] = nums[j]
				nums[j] = temp
				current++
			}
		}
		iter++
		modnum *= 10
	}
	for i := 0; i < 10; i++ {
		fmt.Println(nums2[i])
	}
}
