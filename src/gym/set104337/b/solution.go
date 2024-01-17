package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		l, r := readTwoNums(reader)
		res := solve(l, r)
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

func solve(l int, r int) int {
	res := calc(r)
	if l > 0 {
		res -= calc(l - 1)
	}
	return res
}

type node [19]int

func calc(num int) int {
	var arr []int
	for tmp := num; tmp > 0; tmp /= 10 {
		arr = append(arr, tmp%10)
	}

	n := len(arr)

	if n == 19 {
		return 18 + calc(num-1)
	}

	dp := make([]map[node]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make(map[node]int)
	}

	freq := make([]int, 10)

	// 但是leading 0 的状态， 怎么处理呢？
	var dfs func(lz bool, eq bool, i int, cur node) int
	dfs = func(lz bool, eq bool, i int, cur node) int {
		if i < 0 {
			if lz {
				return 1
			}
			return getMode(cur)
		}

		var res int
		if !eq {
			if v, ok := dp[i][cur]; ok {
				return v
			}
			defer func() {
				dp[i][cur] = res
			}()
		}
		up := 9
		if eq {
			up = arr[i]
		}

		for x := 0; x <= up; x++ {
			if lz && x == 0 {
				res += dfs(lz, false, i-1, cur)
			} else if cur[freq[x]] > 0 {
				cur[freq[x]]--
				freq[x]++
				cur[freq[x]]++
				res += dfs(false, eq && x == arr[i], i-1, cur)
				cur[freq[x]]--
				freq[x]--
				cur[freq[x]]++
			}
		}

		return res
	}

	return dfs(true, true, n-1, [19]int{10, 0})
}

func getMode(n node) int {
	for i := 18; i >= 0; i-- {
		if n[i] > 0 {
			return i
		}
	}
	return 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
