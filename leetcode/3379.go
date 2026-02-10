package leetcode

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		rem := (i + nums[i]) % n
		if rem < 0 {
			rem += n
		}
		res[i] = nums[rem%n]
	}

	return res
}
