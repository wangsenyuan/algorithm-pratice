package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	boxes := make([][]int, n)
	for i := 0; i < n; i++ {
		boxes[i] = readNNums(reader, 3)
	}
	res := solve(boxes)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const inf = 1 << 30

func solve(boxes [][]int) bool {
	// a[0] < b[0], a[1] < b[1], a[2] < b[2]
	for _, cur := range boxes {
		sort.Ints(cur)
	}

	// 3组关系，每组三种关系,a[0] </>/= b[0]
	// 如果能保证在0/1/2维度，不存在重复的值,
	// 超过27个，必然存在一对，满足条件
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i][0] < boxes[j][0]
	})

	// 记录在位置i处 a[1] < b[1]时最小的a[2]是多少
	xs := make(map[int]int)
	for _, box := range boxes {
		xs[box[1]]++
	}

	arr := make([]int, 0, len(xs))
	for x := range xs {
		arr = append(arr, x)
	}
	sort.Ints(arr)
	for i, x := range arr {
		xs[x] = i
	}
	m := len(xs)
	arr = make([]int, 2*m)

	for i := m; i < 2*m; i++ {
		arr[i] = inf
	}
	for i := m - 1; i > 0; i-- {
		arr[i] = min(arr[i*2], arr[i*2+1])
	}

	set := func(p int, v int) {
		p += m
		arr[p] = min(arr[p], v)
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}
	get := func(l int, r int) int {
		l += m
		r += m
		res := inf
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	for i := 0; i < len(boxes); {
		j := i
		for i < len(boxes) && boxes[i][0] == boxes[j][0] {
			// avoid same height conflicts
			w := xs[boxes[i][1]]
			d := get(0, w)
			if d < boxes[i][2] {
				return true
			}
			i++
		}
		for j < i {
			w := xs[boxes[j][1]]
			d := boxes[j][2]
			set(w, d)
			j++
		}
	}
	return false
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
