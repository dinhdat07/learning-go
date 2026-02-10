package leetcode

func isTrionic(nums []int) bool {
	flag, count := 1, 0

	for i := 1; i < len(nums); i++ {
		sign := nums[i] - nums[i-1]
		if sign*flag > 0 {
			continue
		}

		if i == 1 || count == 2 || sign == 0 {
			return false
		} else {
			flag *= -1
			count++
		}
	}

	return count == 2
}
