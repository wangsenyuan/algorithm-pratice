package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	D := readNum(reader)

	S := make([]string, D)

	for i := 0; i < D; i++ {
		s, _ := reader.ReadString('\n')
		S[i] = s[:len(s)-1]
	}

	D = anagrams(S)

	S = S[:D]

	r, c := readTwoNums(reader)
	n := readNum(reader)

	grid := make([]string, r)

	for n > 0 {
		n--
		for i := 0; i < r; i++ {
			s, _ := reader.ReadString('\n')
			grid[i] = s[:len(s)-1]
		}
		fmt.Println(solve(D, S, r, c, grid))
	}
}

func anagrams(S []string) int {
	sort.Sort(SS(S))
	var p int
	for i := 1; i <= len(S); i++ {
		if i == len(S) || compare(S[i-1], S[i]) < 0 {
			S[p] = S[i-1]
			p++
		}
	}

	return p
}

type SS []string

func (ss SS) Len() int {
	return len(ss)
}

func compare(x, y string) int {
	a := getDigitCount(x)
	b := getDigitCount(y)
	for k := 0; k < 26; k++ {
		if a[k] < b[k] {
			return -1
		} else if a[k] > b[k] {
			return 1
		}
	}
	return 0
}

func (ss SS) Less(i, j int) bool {
	return compare(ss[i], ss[j]) < 0
}

func (ss SS) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func getDigitCount(s string) []int {
	res := make([]int, 26)
	for i := 0; i < len(s); i++ {
		res[int(s[i]-'A')]++
	}

	return res
}

func solve(D int, S []string, r, c int, grid []string) string {
	cnt := make([]int, 26)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			cnt[int(grid[i][j]-'A')]++
		}
	}
	ii := make([]int, 16)
	var mx int
	var p int
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			ii[p] = i
			p++
			mx = max(mx, cnt[i])
		}
	}
	ii = ii[:p]

	ss := make(map[int]bool)

	tmp := make([]int, 26)

	for i := 0; i < D; i++ {
		for j := 0; j < 26; j++ {
			tmp[j] = 0
		}
		var j int
		for j < len(S[i]) {
			a := int(S[i][j] - 'A')
			tmp[a]++
			if tmp[a] > cnt[a] {
				break
			}
			j++
		}

		if j < len(S[i]) {
			continue
		}
		var b int
		for j := 0; j < p; j++ {
			b = b*(mx+1) + tmp[ii[j]]
		}
		ss[b] = true
	}
	P := pow(mx+1, p)
	dp := make([]int, P)
	// dp[0] = 0, lose position
	for i := 1; i < P; i++ {
		dp[i] = -1
	}
	var dfs func(x int) bool

	dfs = func(x int) bool {
		if x == 0 {
			return false
		}
		if dp[x] >= 0 {
			return dp[x] > 0
		}

		dp[x] = 0

		for y := range ss {
			if canMark(x, y, mx+1) {
				if !dfs(mark(x, y, mx+1)) {
					dp[x] = 1
					return true
				}
			}
		}
		return false
	}

	var X int

	for j := 0; j < p; j++ {
		X = X*(mx+1) + cnt[ii[j]]
	}

	if dfs(X) {
		return "Alice"
	}
	return "Bob"
}

func canMark(x, y, b int) bool {
	for y > 0 {
		i := x % b
		j := y % b
		if i < j {
			return false
		}
		x /= b
		y /= b
	}
	return true
}

func mark(x, y, b int) int {
	var res int
	B := 1
	for y > 0 {
		i := x % b
		j := y % b
		res = (i-j)*B + res
		B *= b
		x /= b
		y /= b
	}
	for x > 0 {
		i := x % b
		res = i*B + res
		B *= b
		x /= b
	}

	return res
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
