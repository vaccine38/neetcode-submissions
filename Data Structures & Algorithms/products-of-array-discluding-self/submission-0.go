func productExceptSelf(nums []int) []int {
	n := len(nums)
	lr := make([]int, n)
	rl := make([]int, n)
	ans := make([]int, n)
	lr[0] = 1
	rl[n-1] = 1
	for i := 1; i < n; i++ {
		lr[i] = lr[i-1] * nums[i-1]
		rl[n-i-1] = rl[n-i] * nums[n-i]
	}

	for i := range n {
		if i == 0 {
			ans[i] = rl[0]
		} else if i == n -1 {
			ans[i] = lr[n-1]
		} else {
			ans[i] = lr[i] * rl[i]
		}
	}
	return ans
}
