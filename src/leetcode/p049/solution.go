package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	rs := groupAnagrams(strs)
	for _, r := range rs {
		fmt.Printf("%v\n", r)
	}
}

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Sort(asByte(bytes))
		key := string(bytes)
		if as, ok := group[key]; ok {
			group[key] = append(as, str)
		} else {
			group[key] = []string{str}
		}
	}

	rs := make([][]string, 0, len(group))

	for _, v := range group {
		rs = append(rs, v)
	}
	return rs
}

type asByte []byte

func (a asByte) Len() int {
	return len(a)
}

func (a asByte) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a asByte) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
