package p088

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
		i--
	}

	for n > 0 {
		nums1[i] = nums2[n-1]
		i--
		n--
	}
}
