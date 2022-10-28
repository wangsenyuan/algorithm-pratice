package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(l, r int) int {
		l++
		r++
		fmt.Printf("1 %d %d\n", l, r)
		return readNum(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)

		l, r := solve(n, k, ask)
		fmt.Printf("2 %d %d\n", l, r)
		readNum(reader)
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, k int, ask func(int, int) int) (int, int) {
	// find left-most position where A[i] = 1
	// some place have value 1
	p := search(0, n, func(i int) bool {
		return ask(0, i) < 2*(i+1)
	})

	if p == n {
		return 1, k / 2
	}

	// A[p] = 1, and A[:p] = 2
	suf := ask(p, n-1)

	if suf == k {
		return p + 1, n
	}

	if suf > k {
		r := search(p, n, func(i int) bool {
			return ask(p, i) > k
		})
		r--
		x := ask(p, r)
		// x <= k
		if x == k {
			return p + 1, r + 1
		}
		return p + 2, r + 2
	}
	// tmp < k
	if suf&1 == k&1 {
		x := (k - suf) / 2
		return p - x + 1, n
	}
	// suf is sum now
	suf += 2 * p
	// need find right-most 1 in the array
	q := search(0, n, func(i int) bool {
		return ask(i, n-1) == 2*(n-i)
	})
	q--
	// q is last position of 1
	suf -= 2*(n-1-q) + 1
	return (suf-k)/2 + 1, q
}

func search(l, r int, f func(int) bool) int {
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
