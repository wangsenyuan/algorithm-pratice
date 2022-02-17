package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	scanner := bufio.NewReader(os.Stdin)

	tc := readNum(scanner)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, Q := readTwoNums(scanner)
		A := readNNums(scanner, n)
		K := make([]int, Q)
		for i := 0; i < Q; i++ {
			K[i] = readNum(scanner)
		}
		R := solve(n, Q, A, K)
		for _, r := range R {
			buf.WriteString(fmt.Sprintf("%d\n", r))
		}
	}
	fmt.Print(buf.String())
}

func solve(n, Q int, A []int, K []int) []int {
	dp := make([]int, 32)
	fp := make([]int, 32)

	var dfs func(i int, nums []int)

	dfs = func(i int, nums []int) {
		if i < 0 || len(nums) == 0 {
			return
		}
		ones := make([]int, 0, len(nums))
		zeros := make([]int, 0, len(nums))

		for j := 0; j < len(nums); j++ {
			num := nums[j]
			x := (num >> uint(i)) & 1
			if x == 1 {
				ones = append(ones, num)
				dp[i] += len(zeros)
			} else {
				zeros = append(zeros, num)
				fp[i] += len(ones)
			}
		}
		dfs(i-1, ones)
		dfs(i-1, zeros)
	}

	dfs(31, A)

	R := make([]int, Q)

	for q := 0; q < Q; q++ {
		k := K[q]
		var r int

		for i := 31; i >= 0; i-- {
			x := (k >> uint(i)) & 1
			if x == 0 {
				// will not change the result
				r += fp[i]
			} else {
				// will change
				r += dp[i]
			}
		}

		R[q] = r
	}

	return R
}
