package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	var buf bytes.Buffer

	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	for tc > 0 {
		tc--
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

type pair struct {
	first  int
	second int
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

func process(reader *bufio.Reader) []int {
	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	d := readNNums(reader, n)
	l := readNNums(reader, m)
	q := readNNums(reader, k)
	return solve(a, d, l, q)
}

const inf = 1 << 60

func solve(a []int, d []int, l []int, q []int) []int {

	// 有些任务不能在d-a处，如果它开始的太晚，有可能造成后续的任务完不成
	// 所以必须要知道每个任务的最晚开始时间
	// x[i] = min(d[i] - a[i], x[j] - a[i]) = min(d[i], x[j]) - a[i]
	// j是在x的后面发生的任务
	type task struct {
		id int
		a  int
		d  int
		st int
	}

	n := len(a)
	tasks := make([]task, n)
	for i := range n {
		tasks[i] = task{i, a[i], d[i], d[i] - a[i]}
	}

	slices.SortFunc(tasks, func(x, y task) int {
		if x.d != y.d {
			return x.d - y.d
		}
		return x.a - y.a
	})

	mst := inf
	for i := n - 1; i >= 0; i-- {
		cur := tasks[i]
		cur.st = min(cur.st, mst-cur.a)
		tasks[i] = cur
		mst = cur.st
	}

	// 现在它们的时间，肯定不一样
	slices.SortFunc(tasks, func(x, y task) int {
		return x.st - y.st
	})

	m := len(l)
	k := len(q)

	qs := make([]pair, k)
	for i, t := range q {
		qs[i] = pair{t, i}
	}

	slices.SortFunc(qs, func(x, y pair) int {
		return x.first - y.first
	})

	var cur int

	ans := make([]int, k)
	var j, u int
	for i := 0; i <= n; i++ {
		gap := inf
		if i < n {
			gap = tasks[i].st - cur
		}
		var sum int
		for j < m && sum+l[j] <= gap {
			sum += l[j]
			for u < k && qs[u].first < cur+sum {
				ans[qs[u].second] = j
				u++
			}

			j++
		}
		cur += sum
		// 必须完成homework
		if i < n {
			cur += tasks[i].a
		}
	}

	for u < k {
		ans[qs[u].second] = m
		u++
	}

	return ans
}
