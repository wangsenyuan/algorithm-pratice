package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {

	res := solve()

	var buf bytes.Buffer

	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d\n", cur))
	}

	fmt.Print(buf.String())
}

const N = 50000
const MAX_L = 11

func solve() []int64 {
	// d(k, 8) >= d(k, 5) >= d(k, 3)
	// and d(k, i) = 0
	// 一个数字是否是 ceil number 非常容易判断
	var res []int64

	process := func(buf []int) {
		res = append(res, toNum(buf))
		n := len(buf)
		for {
			i := n - 2
			for i >= 0 && buf[i] >= buf[i+1] {
				i--
			}
			if i < 0 {
				break
			}
			j := n - 1
			for buf[j] <= buf[i] {
				j--
			}
			buf[i], buf[j] = buf[j], buf[i]
			reverse(buf[i+1:])
			res = append(res, toNum(buf))
		}
	}

	buf := make([]int, MAX_L)

	var k int

	for len(res) < N {
		k++
		// k total digits
		x := (k + 2) / 3
		for i := x; i <= k; i++ {
			for ii := 0; ii < i; ii++ {
				buf[ii] = 8
			}
			y := (k - i + 1) / 2
			// at leat j 5
			for j := y; j <= k-i && j <= i; j++ {
				for jj := 0; jj < j; jj++ {
					buf[i+jj] = 5
				}
				v := k - i - j
				for vv := 0; vv < v; vv++ {
					buf[vv+i+j] = 3
				}
				reverse(buf[:k])
				process(buf[:k])
			}
		}
	}

	sort.Sort(Int64s(res))

	return res[:N]
}

type Int64s []int64

func (this Int64s) Len() int {
	return len(this)
}

func (this Int64s) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64s) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func toNum(arr []int) int64 {
	var res int64

	for i := 0; i < len(arr); i++ {
		res = res*10 + int64(arr[i])
	}
	return res
}
