package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(l int, r int) int {
		fmt.Printf("? %d %d\n", l, r)
		return readNum(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--

		n := readNum(reader)
		res := solve(n, ask)

		fmt.Printf("! %d %d %d\n", res[0], res[1], res[2])
	}
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(n int, ask func(int, int) int) []int {
	A := ask(1, n)
	// len * (len - 1) / 2 + 1 >= value
	ln := search(0, n, func(ln int) bool {
		return ln*(ln-1)/2+1 >= A
	})

	i := search(0, n-ln, func(i int) bool {
		return ask(i+1, n) < A
	})

	i--

	B := ask(i+2, n)
	// A - B = j - i - 1
	j := A - B + i + 1

	pref := (j - i) * (j - i - 1) / 2
	suf := A - pref

	w := search(0, n-j, func(w int) bool {
		return w*(w-1)/2 >= suf
	})

	k := j + w - 1

	return []int{i + 1, j + 1, k + 1}
}

func search(l int, r int, f func(int) bool) int {
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
