package main

import (
	"bufio"
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

	n := readNum(reader)

	query := func(x int) string {
		fmt.Printf("%d\n", x)
		s, _ := reader.ReadString('\n')
		r := s[:1]
		if r == "E" {
			os.Exit(0)
		}
		return r
	}

	solve(n, query)
}

func solve(n int, query func(int) string) {
	rangs := NewRanges(1, n)

	for {
		LN := rangs.Size()

		if LN <= 5 {
			for i := 1; i <= LN; i++ {
				query(rangs.kth(i))
			}
			return
		}

		x := rangs.kth(LN / 2)
		last := query(x)
		L, R := LN/2-1, LN-LN/2
		qr := []int{x, x}

		var done bool

		for !done && L > 1 && R > 1 {
			if last == "G" {
				np := rangs.kth(L * 2 / 3)
				next := query(np)
				if next == last {
					rangs.remove(1, np)
					rangs.remove(qr[0], qr[1])
					done = true
				} else {
					qr[0] = qr[1]
					qr[1] = np
				}
				last = next
				L = rangs.countSmall(np - 1)
			} else {
				np := rangs.kth(LN - R*2/3)
				next := query(np)
				if next == last {
					rangs.remove(np, n)
					rangs.remove(qr[1], qr[0])
					done = true
				} else {
					qr[1] = qr[0]
					qr[0] = np
				}
				last = next
				R = LN - rangs.countSmall(np)
			}
			if !done {
				rangs.remove(min(qr[0], qr[1]), max(qr[0], qr[1]))
				LN = rangs.Size()
				L = rangs.countSmall(min(qr[0], qr[1]))
				R = LN - rangs.countSmall(max(qr[0], qr[1]))
			}
		}

	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type Ranges struct {
	current []Pair
	backup  []Pair
}

func NewRanges(l, r int) *Ranges {
	first := Pair{l, r}
	current := []Pair{first}
	backup := make([]Pair, 0, 1)
	return &Ranges{current, backup}
}

func (this *Ranges) kth(k int) int {
	for _, p := range this.current {
		l := p.second - p.first + 1
		if k > l {
			k -= l
		} else {
			return p.first + k - 1
		}
	}
	return -1
}

func (this Ranges) countSmall(x int) int {
	var res int
	for _, p := range this.current {
		if p.second < x {
			res += p.second - p.first + 1
		} else {
			if x >= p.first {
				res += x - p.first + 1
			}
		}
	}
	return res
}

func (this Ranges) Size() int {
	var res int
	for _, p := range this.current {
		res += p.second - p.first + 1
	}
	return res
}

func (this *Ranges) remove(l, r int) {
	this.backup = this.backup[:0]
	for _, e := range this.current {
		if e.second < l {
			this.backup = append(this.backup, e)
		} else if e.first > r {
			this.backup = append(this.backup, e)
		} else {
			if e.first < l {
				this.backup = append(this.backup, Pair{e.first, l - 1})
			}
			if e.second > r {
				this.backup = append(this.backup, Pair{r + 1, e.second})
			}
		}
	}
	this.current, this.backup = this.backup, this.current
}
