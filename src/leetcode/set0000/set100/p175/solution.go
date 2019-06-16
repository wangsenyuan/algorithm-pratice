package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//nums := []int{3, 30, 34, 5, 9}
	nums := []int{0, 0}
	fmt.Println(largestNumber(nums))
}

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}

	sort.Sort(Strs(strs))

	var buf bytes.Buffer

	for i := len(strs) - 1; i >= 0; i-- {
		buf.WriteString(strs[i])
	}

	res := buf.String()

	if res[0] == '0' {
		return "0"
	}
	return res
}

type Strs []string

func (s Strs) Len() int {
	return len(s)
}

func (s Strs) Less(i, j int) bool {
	return cmp(s[i], s[j])
}

func cmp(a, b string) bool {
	if len(a) == 0 {
		return true
	}

	if len(b) == 0 {
		return false
	}

	i := 0
	for ; i < len(a) && i < len(b); i++ {
		x, y := a[i], b[i]
		if x < y {
			return true
		}
		if x > y {
			return false
		}
	}

	if i < len(a) {
		return cmp(a[i:], b)
	}

	return cmp(a, b[i:])
}

func (s Strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
