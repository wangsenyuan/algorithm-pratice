package p1381

type CustomStack struct {
	n   int
	ptr int
	bit []int
}

func Constructor(maxSize int) CustomStack {
	n := maxSize
	bit := make([]int, n+1)
	return CustomStack{n, 0, bit}
}

func (this *CustomStack) Push(x int) {
	if this.ptr < this.n {
		update(this.bit, this.n, this.ptr, x)
		update(this.bit, this.n, this.ptr+1, -x)
		this.ptr++
	}
}

func (this *CustomStack) Pop() int {
	if this.ptr == 0 {
		return -1
	}

	get := func(p int) int {
		p++
		var res int
		for p > 0 {
			res += this.bit[p]
			p -= p & -p
		}
		return res
	}
	this.ptr--

	res := get(this.ptr)

	update(this.bit, this.n, this.ptr, -res)
	update(this.bit, this.n, this.ptr+1, res)

	return res
}

func update(bit []int, n int, p int, v int) {
	p++
	for p <= n {
		bit[p] += v
		p += p & -p
	}
}

func (this *CustomStack) Increment(k int, val int) {

	update(this.bit, this.n, 0, val)

	k = min(k, this.ptr)

	update(this.bit, this.n, k, -val)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
