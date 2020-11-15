package p5601

type OrderedStream struct {
	ptr  int
	arr  []string
	used []bool
}

func Constructor(n int) OrderedStream {
	arr := make([]string, n+1)
	used := make([]bool, n+1)
	return OrderedStream{1, arr, used}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	var res []string
	this.arr[id] = value
	this.used[id] = true
	for this.ptr < len(this.used) && this.used[this.ptr] {
		res = append(res, this.arr[this.ptr])
		this.ptr++
	}
	return res
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */
