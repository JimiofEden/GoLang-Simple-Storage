package main

import "fmt"

func main() {
	fmt.Println(climbStairs(15))
}

func climbStairs(n int) int {
    if n == 0 {
		return 1
	}
    if n == 1 {
        return 1
    }
	
	//dp :=make([]int, n+1)
    a := 1
    b := 1
    var c int
	for i:=2; i<=n; i++ {
		c = a + b
        a = b
        b = c
	}
	return c
}