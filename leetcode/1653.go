package leetcode

import "math"

func minimumDeletions(s string) int {
	n := len(s)
	aPref := make([]int, n+1)
	bPref := make([]int, n+1)

	for i := 1; i <= n; i++ {
		aPref[i] = aPref[i-1]
		bPref[i] = bPref[i-1]
		if s[i-1] == 'a' {
			aPref[i]++
		} else {
			bPref[i]++
		}
	}

	ans := math.MaxInt
	for i := 0; i <= n; i++ {
		countBA := bPref[i] + aPref[n] - aPref[i]
		ans = min(countBA, ans)
	}

	return ans

}
