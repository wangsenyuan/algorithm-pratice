package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if res < 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString(fmt.Sprintf("YES %d\n", res))
		}
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(A []int) int {
	n := len(A)
	// X > n, if there are more than 2 numbers <= n has same value, then it if impossible
	sort.Ints(A)

	for i := 1; i < n; i++ {
		if A[i] == A[i-1] {
			return -1
		}
	}

	// A[i] % a = v
	// (A[i] - v) = x * a
	// A[i] = v + x * a
	// a * x + v = y
	// v = [1...n]
	// find two number that > n
	var it int
	mem := make(map[int]bool)
	for it < n && A[it] <= n {
		mem[A[it]] = true
		it++
	}

	if it == n {
		return n + 1
	}

	check := func(x int) bool {
		mem2 := make(map[int]bool)
		for i := it; i < n; i++ {
			y := A[i] % x
			if y == 0 || y > n || mem[y] || mem2[y] {
				return false
			}
			mem2[y] = true
		}
		return len(mem)+len(mem2) == n
	}

	for v := 1; v <= n; v++ {
		if !mem[v] {
			a := A[it] - v
			// a % x = 0
			for x := 1; x <= a/x; x++ {
				if a%x == 0 {
					if x > n && check(x) {
						return x
					}
					y := a / x
					if y > n && check(y) {
						return y
					}
				}
			}
		}
	}
	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
