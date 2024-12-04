package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		res := process(reader)
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

func process(reader *bufio.Reader) int {
	line := readNNums(reader, 4)
	n, T, a, b := line[0], line[1], line[2], line[3]
	kind := readNNums(reader, n)
	t := readNNums(reader, n)
	return solve(n, T, a, b, kind, t)
}

func solve(n int, T int, a int, b int, kind []int, t []int) int {

	type pair struct {
		first  int
		second int
	}

	exams := make([]pair, n)

	freq := make([]int, 2)

	for i := range n {
		exams[i] = pair{t[i], kind[i]}
		freq[kind[i]]++
	}

	slices.SortFunc(exams, func(x, y pair) int {
		return x.first - y.first
	})

	// 如果在时刻s离开时，可以将必须的都完成，那么就把此时的最优结果记录下来
	// 然后寻找下一个有效的状态
	// 0 for easy, 1 for hard
	cnt := make([]int, 2)

	get := func(s int) int {
		tmp := cnt[0] + cnt[1]
		x := s - (a*cnt[0] + b*cnt[1])
		u := min(x/a, freq[0]-cnt[0])
		tmp += u
		x -= u * a
		u = min(x/b, freq[1]-cnt[1])
		tmp += u
		return tmp
	}

	// 如果结束，最好在某个任务变成mandantory前结束
	var res int
	for _, cur := range exams {
		// 如果在s时离开
		// 在这个时刻前开始的任务，有没有完成？
		s := cur.first - 1
		// 共有这么多时间，可以给所有的任务
		// 所有的之前的任务，都必须在s前完成
		if a*cnt[0]+b*cnt[1] <= s {
			res = max(res, get(s))
		}
		cnt[cur.second]++
	}

	if a*cnt[0]+b*cnt[1] <= T {
		res = max(res, get(T))
	}

	return res
}
