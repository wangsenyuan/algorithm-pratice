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
		queries[i] = readNNums(reader, 2)
	}

	res := solve(a, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

const MA = 200010
const mod = 1e9 + 7

const inf = 1 << 50

var primes []int
var lpf [MA]int

func init() {
	for i := 2; i < MA; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= MA {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(a []int, queries [][]int) []int {
	// 拆解质因子，记录每个质因子出现的次数，已经最小的次数
	arr := getFactors(a, queries)
	fs := make(map[int]*SegTree)
	n := len(a)
	for _, num := range arr {
		fs[num] = NewSegTree(n)
	}

	for i := 0; i < n; i++ {
		num := a[i]
		for num > 1 {
			x := lpf[num]
			var cnt int
			for num%x == 0 {
				cnt++
				num /= x
			}
			if fs[x] != nil {
				fs[x].Add(i, cnt)
			}
		}
	}

	ans := 1

	for x, tr := range fs {
		v := tr.Get(0, n)
		ans = mul(ans, pow(x, v))
	}

	process := func(i int, num int) int {
		// 这个要加上去才行
		for num > 1 {
			x := lpf[num]
			var add int
			for num%x == 0 {
				add++
				num /= x
			}
			if fs[x] != nil {
				prev_min := fs[x].Get(0, n)
				fs[x].Add(i, add)

				cur_min := fs[x].Get(0, n)
				// prev_min <= cur_min
				ans = mul(ans, pow(x, cur_min-prev_min))
			}
		}

		return ans
	}

	res := make([]int, len(queries))

	for i, cur := range queries {
		res[i] = process(cur[0]-1, cur[1])
	}

	return res
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Add(p int, v int) {
	p += t.sz
	t.arr[p] += v
	for p > 1 {
		t.arr[p>>1] = min(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {

	l += t.sz
	r += t.sz

	res := inf

	for l < r {
		if l&1 == 1 {
			res = min(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func getFactors(a []int, queries [][]int) []int {
	occ := make(map[int]int)

	n := len(a)

	ops := make([][]int, n)
	for _, cur := range queries {
		i, x := cur[0]-1, cur[1]
		ops[i] = append(ops[i], x)
	}

	var arr []int

	pws := make(map[int]int)

	process := func(i int) {
		clear(pws)

		for _, num := range ops[i] {
			for num > 1 {
				x := lpf[num]

				if pws[x] == 0 {
					// x是第一次出现在当前位
					occ[x]++

					if occ[x] == n {
						arr = append(arr, x)
					}
				}

				for num%x == 0 {
					pws[x]++
					num /= x
				}
			}
		}
	}

	for i, num := range a {
		ops[i] = append(ops[i], num)
		process(i)
	}
	// len(arr) 应该不会很大的吧？
	return arr
}
