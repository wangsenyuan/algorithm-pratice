package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var t int
		pos := readInt(s, 0, &t)
		if t == 0 {
			queries[i] = make([]int, 3)
			for j := 1; j < 3; j++ {
				pos = readInt(s, pos+1, &queries[i][j])
			}
		} else {
			queries[i] = []int{1, 0}
			readInt(s, pos+1, &queries[i][1])
		}
	}

	res := solve(a, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const blk_size = 300

func solve(a []int, queries [][]int) [][]int {
	// 操作1 a[p] = x
	// 操作2 find last p jump¬ position
	// 如果不考虑1，那么这个就是union-find操作
	// 假设 i -> j -> k
	// 修改 a[j] -> x后，
	//  i -> j -> l
	// 有个想法是分段，在每一段内，使用union-find, 修改任何一条的时候，restore它所在段的部分
	n := len(a)

	fa := make([]int, n)
	// jp[i] = i在离开本blocks前的跳跃次数
	jp := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		jp[i] = 0
		fa[i] = i
		j := i + a[i]
		if j < n && j/blk_size == i/blk_size {
			jp[i] = jp[j] + 1
			fa[i] = fa[j]
		}
	}

	change := func(p int, x int) {
		l := p / blk_size * blk_size
		r := min(n, l+blk_size)
		// change it
		a[p] = x
		// 原来那些指向位置p的，前面的部分，不影响
		for i := r - 1; i >= l; i-- {
			fa[i] = i
			jp[i] = 0
			j := i + a[i]
			if j < n && j/blk_size == i/blk_size {
				jp[i] = jp[j] + 1
				fa[i] = fa[j]
			}
		}
	}

	query := func(i int) []int {
		var jumps int
		for {
			jumps += jp[i]
			i = fa[i]
			if i+a[i] >= n {
				return []int{i + 1, jumps + 1}
			}
			// go to next blk
			i += a[i]
			jumps++
		}
	}

	var ans [][]int

	for _, cur := range queries {
		if cur[0] == 0 {
			// 0 p x
			change(cur[1]-1, cur[2])
		} else {
			ans = append(ans, query(cur[1]-1))
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
