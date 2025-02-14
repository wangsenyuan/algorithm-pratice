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
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) []int {
	n := len(a)

	s1 := make([]int, n)
	var t1 int
	s2 := make([]int, n)
	var t2 int

	l1 := make([]int, n)
	l2 := make([]int, n)
	for i, x := range a {
		l1[i] = -1
		l2[i] = -1
		for t1 > 0 && a[s1[t1-1]] <= x {
			t1--
		}
		if t1 > 0 {
			l1[i] = s1[t1-1]
		}
		s1[t1] = i
		t1++
		for t2 > 0 && a[s2[t2-1]] < x {
			t2--
		}
		if t2 > 0 {
			l2[i] = s2[t2-1]
		}
		s2[t2] = i
		t2++
	}

	r1 := make([]int, n)
	r2 := make([]int, n)
	t1 = 0
	t2 = 0
	for i := n - 1; i >= 0; i-- {
		r1[i] = n
		r2[i] = n
		for t1 > 0 && a[s1[t1-1]] <= a[i] {
			t1--
		}
		if t1 > 0 {
			r1[i] = s1[t1-1]
		}
		s1[t1] = i
		t1++
		for t2 > 0 && a[s2[t2-1]] < a[i] {
			t2--
		}
		if t2 > 0 {
			r2[i] = s2[t2-1]
		}
		s2[t2] = i
		t2++
	}

	type pair struct {
		first  int
		second int
	}
	sum := make([]int, n+1)
	arr := make([]pair, n)

	for i, x := range a {
		sum[i+1] = sum[i] + x
		arr[i] = pair{x, i}
	}

	sort.Slice(arr, func(i int, j int) bool {
		return arr[i].first > arr[j].first
	})

	ans := make([]int, n)
	for _, cur := range arr {
		i := cur.second

		if r2[i]-l2[i] == 2 {
			ans[i] = a[i]
			continue
		}
		s := sum[r1[i]] - sum[l1[i]+1]
		j := l1[i]
		ans[i] = s
		if j >= 0 && s > a[j] {
			ans[i] = max(ans[i], ans[j])
		}
		j = r1[i]
		if j < n && s > a[j] {
			ans[i] = max(ans[i], ans[j])
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
