package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%.8f\n", x))
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

func process(reader *bufio.Reader) []float64 {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const eps = 1e-7

func solve(a []int) []float64 {
	n := len(a)
	ps := make([]pair, 1, n+1)

	var sum int

	for i, num := range a {
		sum += num
		p := pair{i + 1, sum}
		for {
			sz := len(ps)
			if sz <= 1 || ps[sz-1].sub(ps[sz-2]).det(p.sub(ps[sz-1])) > 0 {
				break
			}
			ps = ps[:sz-1]
		}
		ps = append(ps, p)
	}

	var ans []float64

	for i := 1; i < len(ps); i++ {
		a, b := ps[i-1], ps[i]
		l := b.first - a.first
		avg := float64((b.second - a.second)) / float64(l)
		for l > 0 {
			ans = append(ans, avg)
			l--
		}
	}

	return ans
}

type pair struct {
	first  int
	second int
}

func (this pair) add(that pair) pair {
	return pair{this.first + that.first, this.second + that.second}
}

func (this pair) sub(that pair) pair {
	return pair{this.first - that.first, this.second - that.second}
}

func (this pair) det(that pair) int {
	return this.first*that.second - this.second*that.first
}
