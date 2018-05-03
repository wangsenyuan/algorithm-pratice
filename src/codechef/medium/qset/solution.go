package main

func main() {

}

func solve(n int, S []byte, queries [][]int) []int {
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = int(S[i] - '0')
	}

	rmq := ConstructRMQ(n, A)

	m := len(queries)
	res := make([]int, m)

	m = 0

	for _, query := range queries {
		t, a, b := query[0], query[1], query[2]
		if t == 1 {
			rmq.Update(a-1, b)
		} else {
			_0 := rmq.Query(a-1, b-1, 0) + 1
			_1 := rmq.Query(a-1, b-1, 1)
			_2 := rmq.Query(a-1, b-1, 2)
			res[m] = (_0 * (_0 - 1) / 2) + (_1 * (_1 - 1) / 2) + (_2 * (_2 - 1) / 2)

			m++
		}
	}
	return res[:m]
}

type RMQ struct {
	sum  []int
	rem  [][]int
	size int
}

func ConstructRMQ(n int, arr []int) *RMQ {
	rem := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rem[i] = make([]int, 4*n)
	}

	sum := make([]int, 4*n)

	var init func(i int, left int, right int)

	init = func(i int, left int, right int) {
		if left == right {
			sum[i] = arr[left]
			rem[sum[left]%3][i] = 1
			return
		}

		mid := left + (right-left)/2
		init(2*i+1, left, mid)
		init(2*i+2, mid+1, right)

		sum[i] = sum[2*i+1] + sum[2*i+2]
		for k := 0; k < 3; k++ {
			rem[k][i] = rem[k][2*i+1] + rem[k][2*i+2]
		}
	}

	init(0, 0, n-1)
	return &RMQ{sum, rem, n}
}

func (rmq *RMQ) Update(pos int, x int) {
	var process func(i int, left int, right int)
	process = func(i int, left int, right int) {
		if left == right {
			for k := 0; k < 3; k++ {
				rmq.rem[k][i] = 0
			}
			rmq.rem[x%3][i] = 1
			return
		}
		mid := left + (right-left)/2
		if pos <= mid {
			process(2*i+1, left, mid)
		} else {
			process(2*i+2, mid+1, right)
		}

		for k := 0; k < 3; k++ {
			rmq.rem[k][i] = rmq.rem[k][2*i+1] + rmq.rem[k][2*i+2]
		}
	}
	process(0, 0, rmq.size-1)
}

func (rmq RMQ) Query(left int, right int, reminder int) int {
	var process func(i int, l, r int) int
	process = func(i, l, r int) int {
		if l > right || r < left {
			return 0
		}

		if left <= l && r <= right {
			return rmq.rem[reminder][i]
		}

		mid := l + (r-l)/2

		return process(2*i+1, left, mid) + process(2*i+2, mid+1, right)
	}

	return process(0, 0, rmq.size-1)
}
