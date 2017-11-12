package solution

import "sort"

//Play is used to export maxEnvelopes
func Play(envelops [][]int) int {
	return maxEnvelopes(envelops)
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(byShape(envelopes))

	m := make([]int, len(envelopes)+1, len(envelopes)+1)
	L := 0
	for i, a := range envelopes {
		lo, hi := 1, L
		for lo <= hi {
			mid := (lo + hi) / 2
			b := envelopes[m[mid]]
			if b[0] < a[0] && b[1] < a[1] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		m[lo] = i
		if L < lo {
			L = lo
		}
	}

	return L
}

type byShape [][]int

func (a byShape) Len() int {
	return len(a)
}

func (a byShape) Less(i, j int) bool {
	if a[i][0] != a[j][0] {
		return a[i][0] < a[j][0]
	}

	return a[j][1] < a[i][1]
}

func (a byShape) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
