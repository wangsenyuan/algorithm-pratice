package main

import (
	"bufio"
	"os"
	"fmt"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n, m := readTwoNums(scanner)
	scanner.Scan()
	S := scanner.Bytes()
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(scanner, 3)
	}
	res := solve(n, S, queries)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(n int, S []byte, queries [][]int) []int64 {
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = int(S[i]-'0') % 3
	}

	rmq := ConstructRMQ(n, A)

	m := len(queries)
	res := make([]int64, m)

	m = 0

	for _, query := range queries {
		t, a, b := query[0], query[1], query[2]
		if t == 1 {
			rmq.Update(a-1, b)
		} else {
			res[m] = rmq.Query(a-1, b-1)
			m++
		}
	}
	return res[:m]
}

type Node struct {
	prefix []int
	suffix []int
	cnt    int64
	total  int
}

type RMQ struct {
	tree []*Node
	size int
}

func assign(x int) *Node {
	prefix := make([]int, 3)
	suffix := make([]int, 3)
	prefix[x%3] = 1
	suffix[x%3] = 1
	total := x % 3
	var cnt int64
	if x%3 == 0 {
		cnt = 1
	}
	return &Node{prefix, suffix, cnt, total}
}

func merge(left, right *Node) *Node {
	parent := &Node{}
	parent.cnt = left.cnt + right.cnt
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (i+j)%3 == 0 {
				parent.cnt += int64(left.suffix[i] * right.prefix[j])
			}
		}
	}
	prefix := make([]int, 3)
	suffix := make([]int, 3)
	for i := 0; i < 3; i++ {
		prefix[i] = left.prefix[i] + right.prefix[(6-left.total+i)%3]
		suffix[i] = right.suffix[i] + left.suffix[(6-right.total+i)%3]
	}
	parent.prefix = prefix
	parent.suffix = suffix
	parent.total = (left.total + right.total) % 3
	return parent
}

func ConstructRMQ(n int, arr []int) *RMQ {
	m := 4 * n
	tree := make([]*Node, m)

	var init func(i int, left, right int)

	init = func(i int, left, right int) {
		if left == right {
			tree[i] = assign(arr[left])
			return
		}
		mid := (left + right) / 2
		init(2*i+1, left, mid)
		init(2*i+2, mid+1, right)
		tree[i] = merge(tree[2*i+1], tree[2*i+2])
	}
	init(0, 0, n-1)
	return &RMQ{tree, n}
}

func (rmq *RMQ) Update(pos int, x int) {
	var process func(i int, left int, right int)
	process = func(i int, left int, right int) {
		if left > pos || right < pos {
			return
		}
		if left == right {
			rmq.tree[i] = assign(x)
			return
		}
		mid := (left + right) / 2
		process(2*i+1, left, mid)
		process(2*i+2, mid+1, right)
		rmq.tree[i] = merge(rmq.tree[2*i+1], rmq.tree[2*i+2])
	}
	process(0, 0, rmq.size-1)
}

func (rmq RMQ) Query(start, end int) int64 {
	var loop func(i int, left int, right int) *Node
	loop = func(i int, left int, right int) *Node {
		if left > end || right < start {
			return nil
		}
		if start <= left && right <= end {
			return rmq.tree[i]
		}
		mid := (left + right) / 2
		q1 := loop(2*i+1, left, mid)
		q2 := loop(2*i+2, mid+1, right)
		if q1 == nil {
			return q2
		}
		if q2 == nil {
			return q1
		}
		return merge(q1, q2)
	}
	return loop(0, 0, rmq.size-1).cnt
}
