package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(5, 5))
}

func uniquePaths(m int, n int) int {
	if m == 1 {
		return 1
	} else if n == 1 {
		return 1
	} else {

		var dp = make([][]int, m)
		for j := range dp {
			dp[j] = make([]int, n)
		}

		for i := 0; i < m; i++ {
			dp[i][0] = 1
		}
		for j := 0; j < n; j++ {
			dp[0][j] = 1
		}
		for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
		return dp[m-1][n-1]
	}
}
