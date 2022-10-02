package p2423

type LUPrefix struct {
	arr []int
	sz  int
}

func Constructor(n int) LUPrefix {
	arr := make([]int, n*2)

	for i := n; i < 2*n; i++ {
		arr[i] = i - n
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min(arr[i*2], arr[i*2+1])
	}

	return LUPrefix{arr, n}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (this *LUPrefix) Upload(video int) {
	video--
	video += this.sz
	this.arr[video] = this.sz

	for video > 1 {
		this.arr[video>>1] = min(this.arr[video], this.arr[video^1])
		video >>= 1
	}
}

func (this *LUPrefix) Longest() int {
	l, r := 0, this.sz
	res := this.sz
	l += this.sz
	r += this.sz
	for l < r {
		if l&1 == 1 {
			res = min(res, this.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, this.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */
