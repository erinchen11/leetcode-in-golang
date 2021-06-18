package climbingstairs


func climbStairs(n int) int {

	if n <= 1 {
		return n
	}

	prev, curr := 1, 1
	for i := 0; i < n; i++ {
		prev, curr = curr, prev + curr
	}
	return curr



}