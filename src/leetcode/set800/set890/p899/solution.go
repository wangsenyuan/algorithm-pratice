package p899

import "sort"

func orderlyQueue(S string, K int) string {
	bs := []byte(S)
	if K > 1 {
		sort.Sort(ByteSlice(bs))
		return string(bs)
	}
	n := len(S)
	S = S + S
	var j int
	for i := 1; i < n; i++ {
		X := S[i : i+n]
		if X < S[j:j+n] {
			j = i
		}
	}
	return S[j : j+n]
}

type ByteSlice []byte

func (bs ByteSlice) Len() int {
	return len(bs)
}

func (bs ByteSlice) Less(i, j int) bool {
	return bs[i] < bs[j]
}

func (bs ByteSlice) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
