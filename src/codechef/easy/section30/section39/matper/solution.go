package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, m := readTwoNums(reader)
	S := readString(reader)
	P := make([]string, m)
	for i := 0; i < m; i++ {
		P[i] = readString(reader)
	}
	res := solve(S, P)
	fmt.Println(res)
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

func solve(S string, P []string) int64 {
	// 先计算出P[i]在S中可能出现的所有位置 vec[i]
	// 然后枚举 P的排列 14!, 检查是否成立
	// dp[state]？
	// 以state表示的集合覆盖S的最小的位置？
	n := len(S)
	m := len(P)
	R := reverse(S)
	nxf := make([][]int, m)
	nxr := make([][]int, m)
	L := make([]int, m)

	fillNext := func(nxt []int, s string, p string) {
		lps := kmp(p)
		for i := 0; i < len(nxt); i++ {
			nxt[i] = n
		}
		for i, j := 0, 0; i < n; {
			if s[i] == p[j] {
				i++
				j++
			}
			if j == len(p) {
				nxt[i-j] = i - j
				j = lps[j-1]
			} else if i < n && s[i] != p[j] {
				if j > 0 {
					j = lps[j-1]
				} else {
					i++
				}
			}
		}
		for i := n - 2; i >= 0; i-- {
			nxt[i] = min(nxt[i], nxt[i+1])
		}
	}

	for i := 0; i < m; i++ {
		L[i] = len(P[i])
		nxf[i] = make([]int, n)
		nxr[i] = make([]int, n)
		fillNext(nxf[i], S, P[i])
		fillNext(nxr[i], R, reverse(P[i]))
	}

	calc := make([]int64, n+1)

	reset := func() {
		for i := 0; i <= n; i++ {
			calc[i] = 0
		}
	}

	var permute func(arr []int, pos int, nxt [][]int, flag bool) int64

	permute = func(arr []int, pos int, nxt [][]int, flag bool) int64 {
		if pos == len(arr) {
			var cur int
			for i := 0; i < len(arr); i++ {
				if cur == n || nxt[arr[i]][cur] == n {
					return 0
				}
				cur = nxt[arr[i]][cur] + L[arr[i]]
			}
			if flag {
				return calc[n-cur]
			}
			calc[cur]++
			return 0
		}
		var res int64
		for i := pos; i < len(arr); i++ {
			arr[pos], arr[i] = arr[i], arr[pos]
			res += permute(arr, pos+1, nxt, flag)
			arr[pos], arr[i] = arr[i], arr[pos]
		}
		return res
	}
	left := make([]int, m/2)
	right := make([]int, m-m/2)
	var ans int64
	for mask := 0; mask < 1<<m; mask++ {
		cnt := digitCount(mask)
		if cnt != m/2 {
			continue
		}

		var j int

		for i := 0; i < m; i++ {
			if (mask>>i)&1 == 1 {
				left[j] = i
				j++
			}
		}
		reset()

		permute(left, 0, nxf, false)

		for i := 1; i <= n; i++ {
			calc[i] += calc[i-1]
		}
		j = 0
		for i := 0; i < m; i++ {
			if (mask>>i)&1 == 0 {
				right[j] = i
				j++
			}
		}

		ans += permute(right, 0, nxr, true)
	}
	return ans
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func kmp(s string) []int {
	n := len(s)
	lps := make([]int, n)

	for i := 1; i < n; i++ {
		j := lps[i-1]

		for j > 0 && s[i] != s[j] {
			j = lps[j-1]
		}

		lps[i] = j
		if s[i] == s[j] {
			lps[i]++
		}
	}

	return lps
}
