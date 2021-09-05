package main

import "fmt"

func main() {
	cost := []int{4, 5, 2, 1, 6, 4, 2, 3}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
    return getMinStairCost(len(cost), 2, cost, true)
}

func getMinStairCost(n int, k int, ps []int, extraTopStair bool) int {
	if n == 0 {
		return 0
	}
	if k == 1 {
		return ps[1]
	}

	if (extraTopStair) {
		ps = append(ps, 0)
		n++
	}

	arraySum := 0
	for i:=0; i<len(ps);i++ {
		arraySum += ps[i]
	}
	
	dp := make([]int, n+1)
	dp[0] = 0
	for i:=1; i<=n; i++ {
		kArr := make([]int, k)
		for j:=1; j<=k; j++ {

			if (i-j >= 0) {
				kArr[j-1] += ps[i-1] + dp[i-j]
			} else {
				kArr[j-1] += arraySum
			}
		}

		min := kArr[0]
		for _, value := range kArr {
			if (value < min) {
				min = value
			}
		}

		dp[i] += min
	}
	return dp[n]
}