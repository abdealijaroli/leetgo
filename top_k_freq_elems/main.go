package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{}

	for _, num := range nums {
		countMap[num]++
	}

	countSlice := make([][]int, len(nums)+1)

	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	res := make([]int, 0, k)

	for i := len(countSlice) - 1; i > 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			break
		}
	}

	return res
}

func main() {
	fmt.Println(topKFrequent([]int{-1, -1}, 1))
}
