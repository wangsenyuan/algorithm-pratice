package p1385

import "sort"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// sort.Ints(arr1)
	sort.Ints(arr2)

	var res int

	for i := 0; i < len(arr1); i++ {
		j := sort.SearchInts(arr2, arr1[i]+d)
		k := sort.SearchInts(arr2, arr1[i]-d)
		if k < j || j < len(arr2) && arr2[j] == arr1[i]+d || k < len(arr2) && arr2[k] == arr1[i]-d {
			continue
		}
		res++
	}

	return res
}
