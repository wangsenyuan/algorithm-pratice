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
		S, _ := reader.ReadString('\n')
		res := solve(n, S)
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

const INF = 1000000010

func solve(n int, S string) int {
	var ok = true
	for i := 1; i < n; i++ {
		if S[i] != S[0] {
			ok = false
			break
		}
	}
	if ok {
		return (n + 2) / 3
	}

	buf := make([]byte, n)
	for i := 1; i < n; i++ {
		if S[i] != S[i-1] {
			copy(buf, S[i:])
			copy(buf[n-i:], S)
			break
		}
	}
	var res int
	var cnt = 1
	for i := 1; i <= n; i++ {
		if i == n || buf[i] != buf[i-1] {
			res += cnt / 3
			cnt = 0
		}
		cnt++
	}
	return res
}

func solve1(n int, S string) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			dp[i][j] = INF
		}
	}

	get := func(i int) int {
		if S[i] == 'L' {
			return 0
		}
		return 1
	}

	valid := func(arr []int) bool {
		for i := 1; i+1 < len(arr); i++ {
			var cnt int
			if arr[i-1] == 1 {
				cnt++
			}
			if arr[i+1] == 0 {
				cnt++
			}
			if cnt == 0 || cnt == 2 {
				continue
			}
			// cnt == 1
			if arr[i-1] == 1 {
				if arr[i] != 0 {
					return false
				}
			}
			if arr[i+1] == 0 {
				if arr[i] != 1 {
					return false
				}
			}
		}
		return true
	}

	valid1 := func(a, b, c int) bool {
		return valid([]int{a, b, c})
	}

	valid2 := func(a, b, c, d int) bool {
		// no RRR no LLL,
		arr := []int{a, b, c, d}
		return valid(arr)
	}

	var best = INF
	for j := 0; j < 8; j++ {
		a, b, c := (j>>2)&1, (j>>1)&1, j&1
		cnt := (a ^ get(0)) + (b ^ get(1)) + (c ^ get(2))

		if valid1(a, b, c) {
			dp[2][j] = cnt
		}
		best = min(best, dp[2][j])
	}

	if n == 3 {
		return best
	}

	best = INF

	for cur := 0; cur < 8; cur++ {
		d0, d1, d2 := (cur>>2)&1, (cur>>1)&1, cur&1
		if !valid1(d0, d1, d2) {
			continue
		}
		for i := 0; i < n; i++ {
			for j := 0; j < 8; j++ {
				dp[i][j] = INF
			}
		}
		cnt := (d0 ^ get(0)) + (d1 ^ get(1)) + (d2 ^ get(2))
		dp[2][cur] = cnt

		for i := 2; i < n-1; i++ {
			for j := 0; j < 8; j++ {
				if dp[i][j] == INF {
					continue
				}
				// valid
				a, b, c := (j>>2)&1, (j>>1)&1, j&1
				for d := 0; d < 2; d++ {
					jj := (j<<1)&7 + d
					if valid2(a, b, c, d) {
						dp[i+1][jj] = min(dp[i+1][jj], dp[i][j]+(get(i+1)^d))
					}
				}
			}
		}
		for j := 0; j < 8; j++ {
			a, b, c := (j>>2)&1, (j>>1)&1, j&1
			if valid2(a, b, c, d0) && valid2(b, c, d0, d1) && valid2(c, d0, d1, d2) {
				best = min(best, dp[n-1][j])
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
