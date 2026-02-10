package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestBalanced(nums []int) int {
	ans, n := 0, len(nums)

	for i := 0; i < n; i++ {
		oddMap, evenMap := make(map[int]struct{}), make(map[int]struct{})

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				evenMap[nums[j]] = struct{}{}
			} else {
				oddMap[nums[j]] = struct{}{}
			}

			if len(oddMap) == len(evenMap) {
				ans = max(ans, j-i+1)
			}
		}

	}

	return ans
}
