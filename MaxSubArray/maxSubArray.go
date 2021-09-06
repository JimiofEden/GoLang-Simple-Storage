package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{0}))
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

func maxSubArray(nums []int) int {
	runningSum := 0
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		runningSum = 0
		for j := i; j < len(nums); j++ {
			runningSum += nums[j]
			if runningSum > maxSum {
				maxSum = runningSum
			}
		}
	}

	return maxSum
}
