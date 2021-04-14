package p179

import (
	"bytes"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	xs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		xs[i] = strconv.Itoa(nums[i])
	}

	sort.Sort(Nums(xs))
	var buf bytes.Buffer

	for i := 0; i < len(xs); i++ {
		buf.WriteString(xs[i])
	}

	res := buf.String()

	for i := 0; i < len(res); i++ {
		if res[i] != '0' {
			return res[i:]
		}
	}
	return "0"
}

type Nums []string

func (ns Nums) Len() int {
	return len(ns)
}

func (ns Nums) Less(i, j int) bool {
	a := ns[i]
	b := ns[j]
	x := a + b
	y := b + a
	var k int
	for k < len(x) && k < len(y) && x[k] == y[k] {
		k++
	}

	return k == len(x) || x[k] > y[k]
}

func (ns Nums) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}
