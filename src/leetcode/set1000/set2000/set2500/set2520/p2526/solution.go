package p2526

type DataStream struct {
	value int
	k     int
	que   []int
	front int
	end   int
	cnt   map[int]int
}

func Constructor(value int, k int) DataStream {
	que := make([]int, k+1)
	front := 0
	end := 0
	cnt := make(map[int]int)
	return DataStream{value, k, que, front, end, cnt}
}

func (this *DataStream) Consec(num int) bool {

	this.que[this.end] = num
	this.end = (this.end + 1) % len(this.que)

	this.cnt[num]++

	sz := (this.end + len(this.que) - this.front) % len(this.que)

	var res bool

	if sz == this.k {
		if len(this.cnt) == 1 {
			for k := range this.cnt {
				res = k == this.value
			}
		}
		this.cnt[this.que[this.front]]--
		if this.cnt[this.que[this.front]] == 0 {
			delete(this.cnt, this.que[this.front])
		}
		this.front = (this.front + 1) % len(this.que)
	}

	return res
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */
