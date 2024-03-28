package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	variables := make([]string, n)
	for i := 0; i < n; i++ {
		variables[i] = readString(reader)
	}

	a, b := solve(m, variables)
	fmt.Println(a)
	fmt.Println(b)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(m int, variables []string) (string, string) {
	n := len(variables)
	vs := make([][]string, n)
	pos := make(map[string]int)
	pos["?"] = n

	vals := make([]string, n)

	for i, cur := range variables {
		vs[i] = parse(cur)
		pos[vs[i][0]] = i
		if len(vs[i]) == 2 {
			vals[i] = determineValue(vs[i][1])
		}
	}

	// 然后从?开始，看它的0/1会造成多少个0/1
	// 但这里有个问题， 0, 导致 a 为0， 但是1也有可能导致 a 为0
	// 但是显然，没法每个bit都尝试后一遍0/1
	// 考虑and， a = b and c, 那么b和c都要是1，才行
	// a = b or c, 要们b和c只要一个是1就行
	// a = b xor c, 那么b和c必须不一致才行
	// 为了获取最小值，当前位最好为0（也不一定）
	// dp[i][j][0/1]表示？第i位位0/1时的情况，当前位是什么
	// 这样子可以消除指数级的计算
	dp := make([][]int, n+1)

	op := func(a, b, c int, o func(a, b int) int) {
		for i := 0; i < 2; i++ {
			dp[a][i] = o(dp[b][i], dp[c][i])
		}
	}

	and := func(a, b, c int) {
		op(a, b, c, func(x, y int) int {
			return x & y
		})
	}

	or := func(a, b, c int) {
		op(a, b, c, func(x, y int) int {
			return x | y
		})
	}

	xor := func(a, b, c int) {
		op(a, b, c, func(x, y int) int {
			return x ^ y
		})
	}

	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if dp[i] != nil {
			return
		}
		dp[i] = make([]int, 2)

		if i == n {
			dp[i][0] = 0
			dp[i][1] = 1
			return
		}

		if vals[i] != "" {
			// 不管?是什么，对i没有影响
			dp[i][0] = int(vals[i][j] - '0')
			dp[i][1] = int(vals[i][j] - '0')
			return
		}
		v := vs[i]
		if len(v) == 2 {
			vi := pos[v[1]]
			dfs(vi, j)
			copy(dp[i], dp[vi])
		} else {
			f, o, s := pos[v[1]], v[2], pos[v[3]]
			dfs(f, j)
			dfs(s, j)
			if o == "AND" {
				and(i, f, s)
			} else if o == "OR" {
				or(i, f, s)
			} else {
				xor(i, f, s)
			}
		}
	}
	// ans[0] for min and ans[1] for max
	ans := make([][]byte, 2)
	for i := 0; i < 2; i++ {
		ans[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			ans[i][j] = '0'
		}
	}

	cnt := make([]int, 2)
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			// reset
			dp[i] = nil
		}
		for i := 0; i < n; i++ {
			if dp[i] == nil {
				dfs(i, j)
			}
		}
		// 当?[j] = 0时，有多少个1
		for x := 0; x < 2; x++ {
			cnt[x] = 0
			for i := 0; i < n; i++ {
				cnt[x] += dp[i][x]
			}
		}
		if cnt[0] < cnt[1] {
			ans[1][j] = '1'
		}
		if cnt[1] < cnt[0] {
			ans[0][j] = '1'
		}
	}

	return string(ans[0]), string(ans[1])
}

func determineValue(s string) string {
	if s[0] >= 'a' && s[0] <= 'z' {
		return ""
	}
	return s
}

func parse(v string) []string {
	i := strings.Index(v, ":=")
	name := v[:i-1]
	v = v[i+3:]
	j := strings.Index(v, " ")
	if j < 0 {
		// a = b
		return []string{name, v}
	}
	first := v[:j]
	v = v[j+1:]
	j = strings.Index(v, " ")
	operand := v[:j]
	second := v[j+1:]
	// name = first operand second
	return []string{name, first, operand, second}
}
