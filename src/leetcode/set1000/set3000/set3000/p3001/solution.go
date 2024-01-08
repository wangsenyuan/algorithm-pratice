package p3001

func maximumSetSize(nums1 []int, nums2 []int) int {
	a := make(map[int]int)
	for _, num := range nums1 {
		a[num]++
	}
	b := make(map[int]int)
	for _, num := range nums2 {
		b[num]++
	}
	n := len(nums1)
	// a的一半 + b的一半
	cnt := []int{n / 2, n / 2}

	set := make(map[int]int)

	for _, num := range nums1 {
		if b[num] == 0 && set[num] == 0 {
			set[num]++
			cnt[0]--
		}
		if cnt[0] == 0 {
			break
		}
	}

	for _, num := range nums2 {
		if a[num] == 0 && set[num] == 0 {
			set[num]++
			cnt[1]--
		}
		if cnt[1] == 0 {
			break
		}
	}

	for k := range set {
		delete(a, k)
		delete(b, k)
	}

	// 剩下的，都是两边都有出现的
	for (cnt[0] > 0 || cnt[1] > 0) && (len(a) > 0 || len(b) > 0) {
		if len(b) == 0 || len(a) > 0 && cnt[0] >= cnt[1] {
			for k := range a {
				delete(a, k)
				if set[k] == 0 && cnt[0] > 0 {
					set[k]++
					cnt[0]--
					break
				}
			}
		} else {
			for k := range b {
				delete(b, k)
				if set[k] == 0 && cnt[1] > 0 {
					set[k]++
					cnt[1]--
					break
				}
			}
		}
	}

	return len(set)
}
