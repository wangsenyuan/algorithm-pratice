package p1577

func numTriplets(nums1 []int, nums2 []int) int {
	return count(nums1, nums2) + count(nums2, nums1)
}

const MAX = 100000

func count(nums1 []int, nums2 []int) int {
	cnt := make(map[int]int)

	for _, num := range nums2 {
		cnt[num]++
	}
	var res int
	for _, num := range nums1 {
		tmp := int64(num) * int64(num)

		for a := range cnt {
			if tmp%int64(a) != 0 {
				continue
			}
			B := tmp / int64(a)
			if B > MAX {
				continue
			}
			b := int(B)

			if a < b {
				res += cnt[a] * cnt[b]
			} else if a == b {
				res += cnt[a] * (cnt[a] - 1) / 2
			}
		}
	}

	return res
}
