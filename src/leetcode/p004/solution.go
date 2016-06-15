package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	k := len(nums1) + len(nums2)
	if k%2 == 0 {
		return (findMedian(nums1, nums2, k-k/2) + findMedian(nums1, nums2, k-(k/2+1))) / 2
	}

	return findMedian(nums1, nums2, k-(k/2+1))
}

func findMedian(as, bs []int, k int) float64 {
	ha := len(as)
	hb := len(bs)

	for k > 0 && ha > 0 && hb > 0 {
		ca := max(1, min(k/2, ha/2))
		cb := max(1, min(k/2, hb/2))
		ma := ha - ca
		mb := hb - cb
		if as[ma] < bs[mb] {
			hb = mb
			k -= cb
		} else {
			ha = ma
			k -= ca
		}
	}

	if ha == 0 && hb == 0 {
		return 0.0
	}

	if ha == 0 {
		return float64(bs[hb-1-k])
	}

	if hb == 0 {
		return float64(as[ha-1-k])
	}

	return float64(max(as[ha-1], bs[hb-1]))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{}
	nums2 := []int{1}

	fmt.Printf("%f\n", findMedianSortedArrays(nums1, nums2))
}
