package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
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

const H = 20

var dp []int

const INF = 1 << 30

func init() {

	dp = make([]int, 1<<H)

	for i := 0; i < len(dp); i++ {
		dp[i] = H
	}
	dp[0] = 0
	for mask := 1; mask < len(dp); mask++ {
		for i := 1; i < H; i++ {
			a := mask >> uint(i)
			b := mask & ((1 << uint(i-1)) - 1)
			dp[mask] = min(dp[mask], dp[a|b]+1)
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve(n int, A []int) int64 {
	AS := make([][]int, H)
	for i := 0; i < H; i++ {
		AS[i] = make([]int, 0, 2)
	}
	B := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		for (1<<uint(tmp))&A[i] == 0 {
			tmp++
		}
		B[i] = tmp - 1
		if tmp > 0 {
			AS[tmp-1] = append(AS[tmp-1], i)
		}
	}

	for i := 0; i < H; i++ {
		AS[i] = append(AS[i], n)
	}

	var res int64

	for i := 0; i < n; i++ {
		var mask int
		if B[i] >= 0 {
			mask = 1 << uint(B[i])
		}
		actPos := i
		for {
			newPos := n
			var bestV int
			for v := 0; v < H; v++ {
				if (mask>>uint(v))&1 == 1 {
					continue
				}
				j := search(len(AS[v]), func(x int) bool {
					return AS[v][x] >= actPos
				})
				canPos := AS[v][j]
				if canPos < newPos {
					newPos = canPos
					bestV = v
				}
			}
			res += int64(newPos-actPos) * int64(dp[mask])
			if newPos == n {
				break
			}
			actPos = newPos
			mask |= 1 << uint(bestV)
		}
	}

	return res
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
