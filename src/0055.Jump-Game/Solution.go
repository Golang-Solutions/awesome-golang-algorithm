package Solution

func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		for j := 0; j <= nums[i] && j < len(nums); j++ {
			if dp[i+j] == 1 {
				dp[i] = 1
				break
			}
		}
	}
	return dp[0] == 1
}

func canJump2(nums []int) bool {
	n, farset := len(nums), 0
	for i := 0; i < n; i++ {
		if farset < i {
			return false
		}
		farset = Max(i+nums[i], farset)
	}
	return true
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
