package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	weapons := make([][]int, n)
	for i := 0; i < n; i++ {
		weapons[i] = readNNums(reader, 2)
	}

	armorSets := make([][]int, m)
	for i := 0; i < m; i++ {
		armorSets[i] = readNNums(reader, 2)
	}

	monsters := make([][]int, k)

	for i := 0; i < k; i++ {
		monsters[i] = readNNums(reader, 3)
	}

	res := solve(weapons, armorSets, monsters)

	fmt.Println(res)
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

func solve(weapons [][]int, armorSets [][]int, monsters [][]int) int {
	weapons = process(weapons)
	// armor sets 好像不需要处理
	armorSets = process(armorSets)

	// 怪物按照防御升序排列
	slices.SortFunc(monsters, func(a, b []int) int {
		return a[0] - b[0]
	})

	//把他们都算在一起，对结果应该影响不大
	vals := make(map[int]int)

	for _, a := range armorSets {
		vals[a[0]]++
	}

	for _, m := range monsters {
		vals[m[1]]++
	}

	arr := make([]int, 0, len(vals))
	for k := range vals {
		arr = append(arr, k)
	}
	sort.Ints(arr)

	for i, k := range arr {
		vals[k] = i
	}

	// 先把b[i]的值更新进去
	arr = make([]int, len(vals))
	for i := 0; i < len(arr); i++ {
		arr[i] = -inf
	}
	for _, a := range armorSets {
		v, c := a[0], a[1]
		v = vals[v]
		arr[v] = -c
	}

	tr := Build(arr)

	n := len(arr)

	best := -inf

	for i, j := 0, 0; i < len(weapons); i++ {
		// 现在要激活weapons[i]
		v, c := weapons[i][0], weapons[i][1]
		for j < len(monsters) && monsters[j][0] < v {
			// 可以被消灭
			w := monsters[j][1]
			w = vals[w]
			z := monsters[j][2]
			tr.Update(w+1, n-1, z)
			j++
		}

		tmp := -c + tr.Get(0, n-1)

		best = max(best, tmp)

	}

	return best
}

func process(arr [][]int) [][]int {
	// 保证 arr[0][0] < arr[1][0]

	slices.SortFunc(arr, func(a, b []int) int {
		if a[1] != b[1] {
			return a[1] - b[1]
		}
		// 相同价格是，威力越大越好
		return b[0] - a[0]
	})

	var res [][]int

	for i := 0; i < len(arr); {
		res = append(res, arr[i])
		j := i
		for i < len(arr) && arr[i][0] <= arr[j][0] {
			// 后面威力小的那些，但是更贵的，全部舍弃掉
			i++
		}
	}

	return res
}

const inf = 1 << 50

type Tree struct {
	vals []int
	lazy []int
	n    int
}

func Build(arr []int) *Tree {
	n := len(arr)
	vals := make([]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			vals[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		vals[i] = max(vals[i*2+1], vals[i*2+2])
	}

	build(0, 0, n-1)

	lazy := make([]int, 4*n)

	return &Tree{vals, lazy, n}
}

func (tr *Tree) update(i int, v int) {
	tr.vals[i] += v
	tr.lazy[i] += v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

// 在区间L...R上增加v
func (tr *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if r < L || R < l {
			return
		}

		if L <= l && r <= R {
			tr.update(i, v)
			return
		}
		tr.push(i)

		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)

		tr.vals[i] = max(tr.vals[i*2+1], tr.vals[i*2+2])
	}

	loop(0, 0, tr.n-1)
}

// 获取区间L...R的最大值
func (tr *Tree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if r < L || R < l {
			return -inf
		}
		if L <= l && r <= R {
			return tr.vals[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		x := loop(2*i+1, l, mid)
		y := loop(2*i+2, mid+1, r)
		return max(x, y)
	}

	return loop(0, 0, tr.n-1)
}
