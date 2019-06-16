package p729

import "sort"

type MyCalendar struct {
	opens []int
	ends  map[int]int
	count int
}

func Constructor() MyCalendar {
	return MyCalendar{make([]int, 1000), make(map[int]int), 0}
}

func (this *MyCalendar) Book(start int, end int) bool {
	n := this.count
	if n == 0 {
		this.opens[n] = start
		this.ends[start] = end
		this.count++
		return true
	}
	i := sort.Search(n, func(i int) bool {
		return this.opens[i] >= start
	})
	if i > 0 {
		before := this.opens[i-1]
		if this.ends[before] > start {
			return false
		}
	}

	if i < this.count {
		if this.opens[i] < end {
			return false
		}

		copy(this.opens[i+1:], this.opens[i:])
		this.opens[i] = start
		this.ends[start] = end
		this.count++
		return true
	}
	this.opens[this.count] = start
	this.ends[start] = end
	this.count++
	return true
}
