package p1428

import "sort"

func checkIfCanBreak(s1 string, s2 string) bool {
	x := []byte(s1)
	y := []byte(s2)

	sort.Sort(Bytes(x))
	sort.Sort(Bytes(y))

	if x[0] >= y[0] {
		var ok = true
		for i := 1; i < len(x); i++ {
			if x[i] < y[i] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	if x[0] <= y[0] {
		var ok = true
		for i := 1; i < len(x); i++ {
			if x[i] > y[i] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	return false
}

type Bytes []byte

func (bs Bytes) Len() int {
	return len(bs)
}

func (bs Bytes) Less(i, j int) bool {
	return bs[i] < bs[j]
}

func (bs Bytes) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
