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

func solve(A []int) int {
	n := len(A)
	sort.Ints(A)

	for i := n/2 - 1; i < n; i++ {
		if A[i-n/2+1] == A[i] {
			return -1
		}
	}

	check := func(g int, pos int) bool {
		var cnt int = 1
		for j := pos + 1; j < n && cnt < n/2; j++ {
			diff := A[j] - A[pos]
			if diff%g == 0 {
				cnt++
			}
		}
		return cnt == n/2
	}

	// 首先可以确定的一点是，A[i] - A[j] 的因数是一个candidate
	// 看能不能在 A[j]前面也找到
	var ans int
	for i := 0; i < n; i++ {
		// A[i] is the base
		for j := i + 1; j < n; j++ {
			if A[j] == A[i] {
				continue
			}
			diff := A[j] - A[i]
			// g is a factor of diff
			for g := 1; g <= diff/g; g++ {
				if diff%g == 0 {
					if check(g, i) {
						ans = max(ans, g)
					}
					if diff/g != g && check(diff/g, i) {
						ans = max(ans, diff/g)
					}
				}
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
