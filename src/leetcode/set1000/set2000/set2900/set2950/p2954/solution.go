package p2954

func findPeaks(mountain []int) []int {
	n := len(mountain)

	var res []int

	for i := 1; i+1 < n; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			res = append(res, i)
		}
	}

	return res
}
