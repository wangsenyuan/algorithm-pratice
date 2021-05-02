package p1845

type SeatManager struct {
	arr []int
}

func Constructor(n int) SeatManager {
	arr := make([]int, n+1)
	return SeatManager{arr}
}

func get(arr []int, p int) int {
	p++
	var res int
	for p > 0 {
		res += arr[p]
		p -= p & -p
	}
	return res
}

func update(arr []int, n int, p int, v int) {
	p++
	for p <= n {
		arr[p] += v
		p += p & -p
	}
}

func (this *SeatManager) Reserve() int {
	n := len(this.arr) - 1
	var l, r = 0, n
	for l < r {
		mid := (l + r) / 2
		if get(this.arr, mid) <= mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	update(this.arr, n, r, 1)
	return r + 1
}

func (this *SeatManager) Unreserve(seatNumber int) {
	update(this.arr, len(this.arr)-1, seatNumber-1, -1)
}
