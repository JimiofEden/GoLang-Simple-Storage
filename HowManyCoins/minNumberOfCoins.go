func coinChange(coins []int, amount int) int {

	return minCoinsToMakeChange(amount, coins)
}

func minCoinsToMakeChange(n int, coinTypes []int) int {

	// Given unliminited number of coins, what is the minimum number required to reach size n?
	// F(i) is the min number of coins at size i)
	// F(i) is the min number of coins at size i with last coin j

	for _, coin := range coinTypes {
		if coin == n {
			return 1
		}
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
		for _, coin := range coinTypes {
			if i-coin < 0 || 1+dp[i-coin] == math.MinInt64 {
				continue
			}
			dp[i] = getMin(dp[i], 1+dp[i-coin])
		}
	}

	if dp[n] == math.MaxInt64 {
		return -1
	}
	return dp[n]
}

func getMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}