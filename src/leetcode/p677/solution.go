package main

import (
	"sort"
	"fmt"
)

type MapSum struct {
	vals map[string]int
	keys []string
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{make(map[string]int), make([]string, 0, 10)}
}

func (this *MapSum) Insert(key string, val int) {
	if _, found := this.vals[key]; found {
		this.vals[key] = val
	} else {
		this.vals[key] = val
		i := sort.SearchStrings(this.keys, key)
		dst := make([]string, len(this.keys)+1)
		copy(dst[:i], this.keys[:i])
		dst[i] = key
		copy(dst[i+1:], this.keys[i:])
		this.keys = dst
	}
}

func (this *MapSum) Sum(prefix string) int {
	i := sort.SearchStrings(this.keys, prefix)
	if i == len(this.keys) {
		return 0
	}

	bs := []byte(prefix)
	n := len(bs)
	bs[n-1] = bs[n-1] + 1
	prefix = string(bs)
	j := sort.SearchStrings(this.keys, prefix)
	ans := 0
	for k := i; k < j; k++ {
		ans += this.vals[this.keys[k]]
	}
	return ans
}

func main() {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	fmt.Println(mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	fmt.Println(mapSum.Sum("ap"))
}
