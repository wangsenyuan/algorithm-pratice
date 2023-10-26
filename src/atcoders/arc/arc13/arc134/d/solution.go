package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, 2*n)

	res := solve(A)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 30

func solve(arr []int) []int {
	n := len(arr) / 2
	// 要获得字典序最小的序列
	//    得到的结果是  a[i] a[j] a[i+n] a[j+n] 而不是 a[i] a[j] a[j+n] a[i+n]
	// 考虑一下，最后的结果形态是什么样子的
	// 假设长度是m, 前半段长m/2
	// 那么前半段肯定是递增的
	//  如果不是，把中间那些超出的部分去掉，会得到更好的结果
	//  后半段的值，至少比前半段它对应i后面的的最大值要大(且不相等）
	//   否则的话，把i后面的部分给删除掉，会得到更优的结果（更小或者更短）
	a := make([]int, n)
	b := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		for p > 0 && a[p-1] > arr[i] {
			p--
		}
		a[p] = arr[i]
		b[p] = arr[i+n]
		p++
	}
	a = a[:p]
	b = b[:p]

	// a is non-dec
	i := bisectRight(a, a[0])
	// a[i] = a[0]
	mn := b[0]
	for j := 0; j < i; j++ {
		mn = min(mn, b[j])
	}
	if mn <= a[0] {
		return []int{a[0], mn}
	}
	// mn > a[0]
	l := bisectLeft(a, b[0])
	r := bisectRight(a, b[0])
	return minOf(makeArray(a[:l], b[:l]), makeArray(a[:r], b[:r]))
}

func bisectLeft(arr []int, x int) int {
	i := search(len(arr), func(i int) bool {
		return arr[i] >= x
	})
	return i
}

func bisectRight(arr []int, x int) int {
	i := search(len(arr), func(i int) bool {
		return arr[i] > x
	})
	return i
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minOf(a []int, b []int) []int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return a
		}
		if a[i] > b[i] {
			return b
		}
	}
	if len(a) <= len(b) {
		return a
	}
	return b
}

func makeArray(a []int, b []int) []int {
	res := make([]int, len(a)+len(b))
	copy(res, a)
	copy(res[len(a):], b)
	return res
}
