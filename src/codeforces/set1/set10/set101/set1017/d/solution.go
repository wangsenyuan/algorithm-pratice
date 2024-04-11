package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, q := readThreeNums(reader)
	w := readNNums(reader, n)
	s := make([]string, m)
	for i := 0; i < m; i++ {
		s[i] = readString(reader)
	}
	x := make([]string, q)
	k := make([]int, q)
	for i := 0; i < q; i++ {
		cur, _ := reader.ReadBytes('\n')
		x[i] = string(cur[:n])
		readInt(cur, n+1, &k[i])
	}

	res := solve(w, s, x, k)

	var buf bytes.Buffer

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const X = 101

func solve(w []int, s []string, x []string, k []int) []int {
	reverse(w)

	n := len(w)
	tot := 1 << n

	cnt := make([]int, tot)
	for _, tmp := range s {
		cnt[convert(tmp)]++
	}

	dp := make([]int, tot)

	for state := 0; state < tot; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				dp[state] += w[i]
			}
		}
	}

	mask := tot - 1

	fp := make([][]int, tot)
	for i := 0; i < tot; i++ {
		fp[i] = make([]int, X)
		for j := 0; j < tot; j++ {
			u := mask ^ i ^ j
			if dp[u] < X {
				fp[i][dp[u]] += cnt[j]
			}
		}
		for j := 1; j < X; j++ {
			fp[i][j] += fp[i][j-1]
		}
	}

	ans := make([]int, len(x))

	for i := 0; i < len(x); i++ {
		u := convert(x[i])
		ans[i] = fp[u][k[i]]
	}

	return ans
}

func solve1(w []int, s []string, x []string, k []int) []int {
	reverse(w)
	n := len(w)
	tot := 1 << n
	cnt := make([]int, tot)
	for _, tmp := range s {
		cnt[convert(tmp)]++
	}

	ans := make([][]int, tot)

	check := func(expect int) []int {
		if len(ans[expect]) > 0 {
			return ans[expect]
		}

		cur := make([]int, X)

		for i := 0; i < tot; i++ {
			if cnt[i] == 0 {
				continue
			}
			var cost int
			for j := 0; j < n; j++ {
				if (i>>j)&1 == (expect>>j)&1 {
					cost += w[j]
				}
				if cost >= X {
					break
				}
			}
			if cost < X {
				cur[cost] += cnt[i]
			}
		}

		for j := 1; j < X; j++ {
			cur[j] += cur[j-1]
		}

		ans[expect] = cur
		return cur
	}

	res := make([]int, len(x))

	for i := 0; i < len(x); i++ {
		v := convert(x[i])
		tmp := check(v)
		res[i] = tmp[k[i]]
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func convert(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res <<= 1
		res |= int(s[i] - '0')
	}
	return res
}
