package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for idx, num := range nums {
		diff := target - num

		if i, ok := m[diff]; ok {
			return []int{i, idx}
		}
		m[num] = idx
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
