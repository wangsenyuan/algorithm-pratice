package main

import (
	"bufio"
	"bytes"
	"os"
	"sort"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) string {
	l, n, m := readThreeNums(reader)
	a := readNNums(reader, l)
	b := make([][]int, n)
	for i := range n {
		b[i] = readNNums(reader, m)
	}
	return solve(a, b)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, b [][]int) string {
	marked := make(map[int]int)
	for _, x := range a {
		marked[x]++
	}
	var arr []int
	for k := range marked {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for i, x := range arr {
		marked[x] = i
	}
	n := len(b)
	m := len(b[0])

	k := len(arr)
	at := make([][]int, k)
	for i := range k {
		at[i] = make([]int, m)
		for j := range m {
			at[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			x := b[i][j]
			if w, ok := marked[x]; ok {
				// 每一行每个数字，只需要记录最后一个位置
				at[w][j] = i
			}
		}
	}
	dp := make([]int, m+1)
	dp[m] = -1
	// 每列里面获胜的行号
	l := len(a)

	fp := make([]int, m+1)
	fp[m] = -1
	// dp是当前的获胜状态
	for l > 0 {
		l--
		x := marked[a[l]]
		if l+1 == len(a) {
			for j := m - 1; j >= 0; j-- {
				dp[j] = max(dp[j+1], at[x][j])
			}
		} else {
			for j := m - 1; j >= 0; j-- {
				r := at[x][j]
				fp[j] = fp[j+1]
				if r >= 0 && dp[j+1] <= r {
					// 这一列有数字x, 且在它的右下角没有获胜的位置
					fp[j] = max(fp[j], r)
				}
			}
			copy(dp, fp)
		}
	}

	if dp[0] >= 0 {
		return "T"
	}
	return "N"
}
