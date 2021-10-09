package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 3)
		}
		res := solve(A, B, Q)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const N = 100010

var bel [N]int
var ff [401][N]int64
var rev [4 * N]int
var wn [4 * N]cplx
var cc [4 * N]cplx
var dd [4 * N]cplx

func solve(A, B []int, Q [][]int) []int64 {
	n := len(A)
	blkSize := 15 * int(math.Sqrt(float64(n)))

	nn := n + blkSize

	l := 1
	lgl := -1
	for l <= nn {
		l <<= 1
		lgl++
	}
	for i := 1; i < l; i++ {
		rev[i] = rev[i>>1]>>1 | (i&1)<<uint(lgl)
		wn[i] = cplx{cos(i, l), sin(i, l)}
	}

	for i := 0; i < l; i++ {
		cc[i].x = 0
		cc[i].y = 0
		dd[i].x = 0
		dd[i].y = 0
	}
	for i := 1; i <= n; i++ {
		cc[i].x = float64(A[i-1])
		bel[i] = (i-1)/blkSize + 1
	}

	dft := func(a []cplx) {
		for i := 0; i < l; i++ {
			if i < rev[i] {
				a[i], a[rev[i]] = a[rev[i]], a[i]
			}
		}
		for i := 1; i < l; i <<= 1 {
			for j := 0; j < l; j += (i << 1) {
				for k := 0; k < i; k++ {
					x := a[j+k]
					y := wn[k*(l/i)].Mul(a[i+j+k])
					a[j+k] = x.Add(y)
					a[j+k+i] = x.Sub(y)
				}
			}
		}
	}

	dft(cc[:N])

	idft := func(a []cplx) {
		for i, j := 1, l-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		dft(a)
		for i := 0; i < l; i++ {
			a[i] = cplx{a[i].x / float64(l), a[i].y}
		}
	}

	for i := 1; i*blkSize <= n; i++ {
		for j := 0; j < l; j++ {
			dd[j].x = 0
			dd[j].y = 0
		}
		for j := 1; j <= blkSize; j++ {
			dd[blkSize+1-j].x = float64(B[(i-1)*blkSize+j-1])
		}
		dft(dd[:N])
		for j := 0; j < l; j++ {
			dd[j] = dd[j].Mul(cc[j])
		}
		idft(dd[:N])
		for j := 1; j <= n; j++ {
			ff[i][j] = int64(dd[j+blkSize].x + 0.5)
		}
	}

	calc := func(x, y, ll int) int64 {
		var res int64
		for i := 0; i < ll; i++ {
			res += int64(A[x+i-1]) * int64(B[y+i-1])
		}
		return res
	}

	ans := make([]int64, len(Q))

	for i, cur := range Q {
		x, y, ll := cur[0], cur[1], cur[2]
		if bel[x] == bel[y] {
			ans[i] = calc(x, y, ll)
			continue
		}
		ans[i] = calc(x, y, bel[y]*blkSize-y+1)
		ans[i] += calc(x+blkSize*(bel[y+ll-1]-1)+1-y, blkSize*(bel[y+ll-1]-1)+1, y+ll-1-blkSize*(bel[y+ll-1]-1))
		for j := bel[y] + 1; j <= bel[y+ll-1]-1; j++ {
			ans[i] += ff[j][x+(j-1)*blkSize+1-y]
		}
	}
	return ans
}

func cos(i, j int) float64 {
	return math.Cos(float64(i) * math.Pi / float64(j))
}

func sin(i, j int) float64 {
	return math.Sin(float64(i) * math.Pi / float64(j))
}

type cplx struct {
	x, y float64
}

func (this cplx) Add(that cplx) cplx {
	return cplx{this.x + that.x, this.y + that.y}
}

func (this cplx) Sub(that cplx) cplx {
	return cplx{this.x - that.x, this.y - that.y}
}

func (this cplx) Mul(that cplx) cplx {
	return cplx{this.x*that.y - this.y*that.x, this.x*that.y + this.y*that.x}
}
