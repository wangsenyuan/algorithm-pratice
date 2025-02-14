package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n, k := readTwoNums(reader)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = readNNums(reader, n)
	}
	return solve(a, k)
}

const mod = 998244353

func mul(a, b int) int {
	return a * b % mod
}

func solve(a [][]int, K int) int {
	n := len(a)

	row := NewDSU(n)
	col := NewDSU(n)

	// 是全部，不是一行

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ok := true
			for k := 0; k < n; k++ {
				if a[i][k]+a[j][k] > K {
					ok = false
					break
				}
			}
			if ok {
				row.Union(i, j)
			}
			ok = true
			for k := 0; k < n; k++ {
				if a[k][i]+a[k][j] > K {
					ok = false
					break
				}
			}
			if ok {
				col.Union(i, j)
			}
		}
	}

	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}

	res := 1
	for i := 0; i < n; i++ {
		if row.Find(i) == i {
			res = mul(res, F[row.cnt[i]])
		}
		if col.Find(i) == i {
			res = mul(res, F[col.cnt[i]])
		}
	}
	return res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
