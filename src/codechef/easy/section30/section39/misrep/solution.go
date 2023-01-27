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
		ok, res := solve(A)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, op := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", op[0], op[1]))
			}
		}
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

func solve(A []int) (bool, [][]int) {
	var sum int
	for _, a := range A {
		sum += a
	}
	if sum%2 != 0 {
		return false, nil
	}
	n := len(A)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true

	for i := 0; i < n; i++ {
		for x := sum; x >= A[i]; x-- {
			dp[i+1][x] = dp[i][x] || dp[i][x-A[i]]
		}
		for x := A[i] - 1; x >= 0; x-- {
			dp[i+1][x] = dp[i][x]
		}
	}
	if !dp[n][sum/2] {
		return false, nil
	}
	var s1 []int
	var s2 []int
	cur := sum / 2
	for i := n - 1; i >= 0; i-- {
		if dp[i+1][cur] && cur >= A[i] && dp[i][cur-A[i]] {
			cur -= A[i]
			s1 = append(s1, i)
		} else {
			s2 = append(s2, i)
		}
	}

	sortIndex := func(arr []int) {
		sort.Slice(arr, func(i, j int) bool {
			return A[arr[i]] > A[arr[j]]
		})
	}

	sortIndex(s1)
	sortIndex(s2)

	var res [][]int

	for i, j := len(s1), len(s2); i > 0 && j > 0; {
		res = append(res, []int{s1[i-1] + 1, s2[j-1] + 1})
		x := min(A[s1[i-1]], A[s2[j-1]])
		A[s1[i-1]] -= x
		A[s2[j-1]] -= x
		if A[s1[i-1]] == 0 {
			i--
		}
		if A[s2[j-1]] == 0 {
			j--
		}
	}
	return true, res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
