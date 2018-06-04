package p846

import "sort"

func isNStraightHand(hand []int, W int) bool {
	n := len(hand)
	if n%W != 0 {
		return false
	}

	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		cnt[hand[i]]++
	}

	group := func(start int) bool {
		for i := 0; i < W; i++ {
			cnt[start+i]--
			if cnt[start+i] < 0 {
				return false
			}
		}
		return true
	}

	sort.Ints(hand)

	for i := 0; i < n; i++ {
		if cnt[hand[i]] > 0 {
			ret := group(hand[i])
			if !ret {
				return false
			}
		}
	}

	return true
}
