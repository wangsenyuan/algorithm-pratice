package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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

func process(reader *bufio.Reader) bool {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const N = 700

var g [N][N]int

var dp [N][N]*BitSet

var fp [N][N][2]bool

func init() {
	for i := range N {
		for j := range N {
			dp[i][j] = NewBitSet(N)
		}
	}
}

func solve(a []int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := range n {
			dp[i][j].Reset()
		}
		dp[i][i].Set(i)
	}

	for i := range n {
		for j := range n {
			g[i][j] = gcd(a[i], a[j])
		}
	}

	for i := range n {
		for j := range n {
			for k := range 2 {
				fp[i][j][k] = false
			}
		}
		if i > 0 && g[i-1][i] > 1 {
			fp[i][i][0] = true
		}
		if i+1 < n && g[i][i+1] > 1 {
			fp[i][i][1] = true
		}
	}

	for r := 0; r < n; r++ {
		for l := r - 1; l >= 0; l-- {
			for i := l; i <= r; i++ {
				if (i == l || fp[l][i-1][1]) && (i == r || fp[i+1][r][0]) {
					dp[l][r].Set(i)
				}
			}
			for i := l; i <= r; i++ {
				if dp[l][r].IsSet(i) {
					if l > 0 && g[l-1][i] > 1 {
						fp[l][r][0] = true
					}
					if r+1 < n && g[i][r+1] > 1 {
						fp[l][r][1] = true
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if dp[0][n-1].IsSet(i) {
			return true
		}
	}
	return false
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type BitSet struct {
	sz  int
	arr []uint64
}

func NewBitSet(n int) *BitSet {
	sz := (n + 63) / 64
	arr := make([]uint64, sz)
	return &BitSet{sz, arr}
}

func (bs *BitSet) Set(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	bs.arr[i] |= 1 << uint64(j)
}

func (bs *BitSet) IsSet(p int) bool {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	return (bs.arr[i]>>uint64(j))&1 == 1
}

func (bs *BitSet) Reset() {
	clear(bs.arr)
}
