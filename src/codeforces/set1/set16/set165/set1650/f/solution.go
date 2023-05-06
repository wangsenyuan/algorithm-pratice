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
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		ops := make([][]int, m)
		for i := 0; i < m; i++ {
			ops[i] = readNNums(reader, 3)
		}
		res := solve(A, ops)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, id := range res {
				buf.WriteString(fmt.Sprintf("%d ", id))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
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

func solve(A []int, options [][]int) []int {
	// A[i] <= A[i+1]
	// option = {ei, ti, pi} 表示在ti时间内，完成任务ei，进度（增加)pi
	// 首先对于所有的任务，如果完成100所需要的time的sum > A[e]，则结果为false
	// 对于第一个任务i， 应该尽可能早的完成它，所以对于任何一个任务，应该计算那些在完成100前提下花费时间最少的options
	// 这个可以用dp来处理

	n := len(A)

	m := len(options)
	ops := make([]Option, m)
	ops_by_task := make([][]Option, n)

	for i := 0; i < m; i++ {
		ops[i] = Option{i + 1, options[i][0] - 1, options[i][1], options[i][2]}
		j := ops[i].task
		if len(ops_by_task[j]) == 0 {
			ops_by_task[j] = make([]Option, 0, 1)
		}
		ops_by_task[j] = append(ops_by_task[j], ops[i])
	}

	dp_buf := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp_buf[i] = make([]int, 101)
	}

	var cur int
	var res []int

	for j := 0; j < n; j++ {
		best_options := findBestOptions(dp_buf, ops_by_task[j], A[j]+1-cur)
		if len(best_options) == 0 {
			// no answer
			return nil
		}
		// use best_operations to finish task[j]
		for _, op := range best_options {
			res = append(res, op.id)
			cur += op.time
		}
		if cur > A[j] {
			return nil
		}
	}

	return res
}

type Option struct {
	id      int
	task    int
	time    int
	percent int
}

func findBestOptions(dp [][]int, ops []Option, inf int) []Option {
	n := len(ops)
	// dp[i] = 在完成任务进度i时，所需要花费的最短时间
	for i := 0; i <= n; i++ {
		for j := 0; j <= 100; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	for i, op := range ops {
		// 不使用option i
		copy(dp[i+1], dp[i])
		for j := 0; j <= 100; j++ {
			if j+op.percent <= 100 {
				dp[i+1][j+op.percent] = min(dp[i+1][j+op.percent], dp[i][j]+op.time)
			} else {
				dp[i+1][100] = min(dp[i+1][100], dp[i][j]+op.time)
			}
		}
		for j := 99; j >= 0; j-- {
			// 完成percent越少，需要的时间越少
			dp[i+1][j] = min(dp[i+1][j], dp[i+1][j+1])
		}
	}

	if dp[n][100] >= inf {
		// no answer
		return nil
	}
	var res []Option
	it := n
	cur := 100
	for ; it > 0 && cur > 0; it-- {
		if it == 1 {
			res = append(res, ops[it-1])
			break
		}
		// it > 0
		if dp[it][cur] == dp[it-1][cur] {
			// can skip it
			continue
		}
		res = append(res, ops[it-1])
		cur -= ops[it-1].percent
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
