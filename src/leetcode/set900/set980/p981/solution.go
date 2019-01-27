package p981

import (
	"sort"
)

type TimeMap struct {
	values     map[string][]string
	timestamps map[string][]int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	values := make(map[string][]string)
	timestamps := make(map[string][]int)
	return TimeMap{values, timestamps}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, found := this.values[key]; !found {
		this.values[key] = make([]string, 0, 10)
		this.timestamps[key] = make([]int, 0, 10)
	}
	this.values[key] = append(this.values[key], value)
	this.timestamps[key] = append(this.timestamps[key], timestamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if vs, found := this.values[key]; !found {
		return ""
	} else {
		ts := this.timestamps[key]
		i := sort.Search(len(ts), func(i int) bool {
			return ts[i] > timestamp
		})
		if i == 0 {
			return ""
		}
		return vs[i-1]
	}
}
