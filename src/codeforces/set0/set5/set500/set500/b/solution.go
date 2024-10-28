package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	p := readNNums(reader, n)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString(reader)
	}
	res := solve(p, a)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
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

func solve(p []int, a []string) []int {
	n := len(p)

	set := NewDSU(n)

	for i, row := range a {
		for j := 0; j < n; j++ {
			if row[j] == '1' {
				set.Union(i, j)
			}
		}
	}

	type pair struct {
		first  int
		second int
	}
	// 同一个集合的可以排序
	arr := make([][]pair, n)
	for i := 0; i < n; i++ {
		j := set.Find(i)
		arr[j] = append(arr[j], pair{p[i], i})
	}
	ans := make([]int, n)
	for _, vs := range arr {
		var pos []int
		for _, v := range vs {
			pos = append(pos, v.second)
		}
		slices.SortFunc(vs, func(x, y pair) int {
			return x.first - y.first
		})
		for i, v := range vs {
			ans[pos[i]] = v.first
		}
	}

	return ans
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
