package p1428

import "sort"

func checkIfCanBreak(s1 string, s2 string) bool {
	x := []byte(s1)
	y := []byte(s2)

	sort.Sort(Bytes(x))
	sort.Sort(Bytes(y))

	return compare(x, y) || compare(y, x)
}

func compare(x, y []byte) bool {
	for i := 0; i < len(x); i++ {
		if x[i] < y[i] {
			return false
		}
	}
	return true
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
