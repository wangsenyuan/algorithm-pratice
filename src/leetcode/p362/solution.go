package main

import (
	"fmt"
	"sort"
)

func main() {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(1)
	hc.Hit(1)
	hc.Hit(300)
	fmt.Printf("%d\n", hc.Gethits(300))
}

//HitCounter is used to record the hits number in the last 5 mins
type HitCounter struct {
	ts  []int
	cnt map[int]int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{make([]int, 0, 100), make(map[int]int)}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	if _, found := this.cnt[timestamp]; !found {
		recordTimestamp(this, timestamp)
	}
	this.cnt[timestamp]++
}

func recordTimestamp(hc *HitCounter, ts int) {
	if len(hc.ts)+1 == cap(hc.ts) {
		tmp := make([]int, len(hc.ts), 2*cap(hc.ts))
		copy(tmp, hc.ts)
		hc.ts = tmp
	}
	hc.ts = append(hc.ts, ts)
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Gethits(timestamp int) int {
	expired, left := sliceWindow(this, timestamp)
	this.ts = left
	for _, ts := range expired {
		delete(this.cnt, ts)
	}

	cnt := 0
	for _, ts := range left {
		if ts > timestamp {
			break
		}
		cnt += this.cnt[ts]
	}
	return cnt
}

func sliceWindow(hc *HitCounter, ts int) ([]int, []int) {
	lo := sort.Search(len(hc.ts), func(i int) bool {
		return hc.ts[i] > ts-300
	})
	return hc.ts[:lo], hc.ts[lo:]
}
