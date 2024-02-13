package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, key := range nums {
		m[key]++
	}

	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys[:k]
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
