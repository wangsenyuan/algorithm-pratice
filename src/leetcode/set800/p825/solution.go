package p825

import "sort"

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	var ans int
	n := len(ages)
	for i := 0; i < n; i++ {
		k := sort.Search(n, func(j int) bool {
			return 2*(ages[j]-7) > ages[i]
		})

		m := sort.Search(n, func(j int) bool {
			return ages[j] > ages[i]
		})

		if m > k {
			ans += m - k
			if k <= i {
				ans--
			}
		}

	}

	return ans
}
