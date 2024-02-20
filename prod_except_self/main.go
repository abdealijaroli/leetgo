package main

import (
	"fmt"
)

// initial solution
func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	post := make([]int, n)

	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		pre[i] = nums[i] * pre[i-1]
	}

	post[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		post[i] = nums[i] * post[i+1]
	}

	res := make([]int, n)

	res[0] = post[0]

	for i := 1; i < n-1; i++ {
		res[i] = pre[i-1] * post[i+1]
	}

	res[n-1] = pre[n-2]

	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
