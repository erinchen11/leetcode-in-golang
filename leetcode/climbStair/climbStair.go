package leetcode

/*

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?


思路：前兩個數相加等於該數
注意：array起始點，要從a[0]=1, a[1]=1開始，所以array起始點為2
*/

func climbStair(n int) int {

	if n == 1 || n == 2 {
		return n
	}

	var nums [1005]int
	nums[1] = 1
	nums[2] = 2

	for i := 3; i <= n; i++ {
		nums[i] = nums[i-2] + nums[i-1]
	}
	return nums[n]

}
