package p2363

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {

	sort.Slice(items1, func(i int, j int) bool {
		return items1[i][0] < items1[j][0]
	})

	sort.Slice(items2, func(i int, j int) bool {
		return items2[i][0] < items2[j][0]
	})

	var res [][]int

	for i, j := 0, 0; i < len(items1) || j < len(items2); {
		if i < len(items1) && j < len(items2) {
			if items1[i][0] < items2[j][0] {
				res = append(res, items1[i])
				i++
				continue
			}
			if items1[i][0] > items2[j][0] {
				res = append(res, items2[j])
				j++
				continue
			}
			res = append(res, []int{items1[i][0], items1[i][1] + items2[j][1]})
			i++
			j++
		} else if j == len(items2) {
			res = append(res, items1[i])
			i++
		} else {
			res = append(res, items2[j])
			j++
		}
	}
	return res
}
