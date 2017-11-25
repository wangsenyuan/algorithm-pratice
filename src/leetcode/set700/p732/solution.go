package p732

import "sort"

type MyCalendarThree struct {
	open  []int
	close []int
	n     int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{make([]int, 1000), make([]int, 1000), 0}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	open, close := this.open, this.close

	i := sort.Search(this.n, func(i int) bool {
		return open[i] > start
	})

	if i < this.n {
		copy(open[i+1:], open[i:])
	}
	open[i] = start

	j := sort.Search(this.n, func(j int) bool {
		return close[j] > end
	})

	if j < this.n {
		copy(close[j+1:], close[j:])
	}
	close[j] = end
	this.n++

	var k, cnt int
	for i, j := 0, 0; i < this.n && j < this.n; {
		if open[i] < close[j] {
			cnt++
			if cnt > k {
				k = cnt
			}
			i++
		} else {
			j++
			cnt--
		}
	}

	return k
}
