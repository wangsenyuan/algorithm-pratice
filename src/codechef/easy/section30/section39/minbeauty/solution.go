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

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(A []int) int {
	sort.Ints(A)
	n := len(A)
	best := 1 << 30
	for j := 1; j < n-1; j++ {
		for i, k := 0, n; i < j; i++ {
			// find A[k] + A[i] - 2 * A[j]
			for k-1 > j && A[k-1]+A[i] >= 2*A[j] {
				k--
			}
			if k > j && k < n {
				tmp := A[k] + A[i] - 2*A[j]
				best = min(best, abs(tmp))
			}
			if k-1 > j {
				tmp := A[k-1] + A[i] - 2*A[j]
				best = min(best, abs(tmp))
			}
		}
	}
	return best
}

func solve1(A []int) int {
	sort.Ints(A)
	n := len(A)
	best := 1 << 30
	for j := 1; j < n-1; j++ {
		for i := 0; i < j; i++ {
			// find A[k] + A[i] - 2 * A[j]
			k := search(j+1, n, func(k int) bool {
				return A[k]+A[i] >= 2*A[j]
			})
			if k < n {
				tmp := A[k] + A[i] - 2*A[j]
				best = min(best, abs(tmp))
			}
			if k-1 > j {
				tmp := A[k-1] + A[i] - 2*A[j]
				best = min(best, abs(tmp))
			}
		}
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
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
