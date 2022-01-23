package p2150

func findLonely(nums []int) []int {
	occ := make(map[int]int)

	for _, num := range nums {
		occ[num]++
	}

	var res []int

	for _, num := range nums {
		cnt := occ[num]
		if cnt > 1 {
			continue
		}
		if _, ok := occ[num+1]; ok {
			continue
		}
		if _, ok := occ[num-1]; ok {
			continue
		}
		res = append(res, num)
	}

	return res
}
