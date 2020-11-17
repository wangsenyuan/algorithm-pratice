package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		s, _ := reader.ReadBytes('\n')
		a, pos := readFrac(s, 0)
		p, pos := readFrac(s, pos+1)
		q, pos := readFrac(s, pos+1)
		var D int
		readInt(s, pos+1, &D)
		res := solve(a, p, q, D)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func readFrac(s []byte, from int) (Frac, int) {
	var res int
	i := from
	for i < len(s) {
		if s[i] == ' ' {
			break
		}
		if s[i] == '.' {
			i++
			continue
		}
		res = res*10 + int(s[i]-'0')
		i++
	}
	return NewFrac(int64(res), 100), i
}

func solve(a, p, q Frac, D int) string {
	d := Frac{1, 1}
	x := p.Mul(a).Mul(d.Sub(a))
	y := q.Mul(a)
	x = x.Add(y)
	z := d.Sub(a)
	z = z.Mul(z)
	ret := x.Div(z)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d.", ret.chi/ret.zna))

	(&ret).chi %= ret.zna

	for i := 1; i <= D; i++ {
		(&ret).chi *= 10
		buf.WriteByte(byte('0' + (ret.chi / ret.zna)))
		(&ret).chi %= ret.zna
	}

	return buf.String()
}

type Frac struct {
	chi, zna int64
}

func NewFrac(a, b int64) Frac {
	c := gcd(a, b)
	return Frac{a / c, b / c}
}

func (this Frac) Add(that Frac) Frac {
	common := this.zna / gcd(this.zna, that.zna) * that.zna
	var res Frac
	res.chi = this.chi*common/this.zna + that.chi*common/that.zna
	res.zna = common
	return res
}

func (this Frac) Sub(that Frac) Frac {
	common := this.zna / gcd(this.zna, that.zna) * that.zna
	var res Frac
	res.chi = this.chi*common/this.zna - that.chi*common/that.zna
	res.zna = common
	return res
}

func (this Frac) Mul(that Frac) Frac {
	var res Frac
	res.chi = this.chi * that.chi
	res.zna = this.zna * that.zna

	g := gcd(res.chi, res.zna)

	return Frac{res.chi / g, res.zna / g}
}

func (this Frac) Div(that Frac) Frac {
	var res Frac
	res.chi = this.chi * that.zna
	res.zna = this.zna * that.chi

	g := gcd(res.chi, res.zna)

	return Frac{res.chi / g, res.zna / g}
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}

	return a
}
