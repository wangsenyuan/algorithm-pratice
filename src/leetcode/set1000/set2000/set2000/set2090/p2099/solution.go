package p2099

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)

	// L[i] = prev index j, that S[j] >= S[j+1] >= S[j+2] >= .. S[i]
	L := make([]int, n)

	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			L[i] = L[i-1]
		} else {
			L[i] = i
		}
	}

	var res []int
	R := n - 1
	for i := n - 1; i >= time; i-- {
		if i+1 < n && security[i] > security[i+1] {
			R = i
		}
		if R-i >= time && i-L[i] >= time {
			res = append(res, i)
		}
	}

	reverse(res)

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
