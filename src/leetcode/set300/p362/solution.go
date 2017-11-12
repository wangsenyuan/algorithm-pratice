package main

import "fmt"

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
	ts []int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{make([]int, 0, 100)}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.ts = append(this.ts, timestamp)
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Gethits(timestamp int) int {
	i := 0
	for i < len(this.ts) && this.ts[i] <= timestamp-300 {
		i++
	}
	this.ts = this.ts[i:]

	return len(this.ts)
}
