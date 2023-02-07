package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		m, n, k := readThreeNums(reader)
		a := readString(reader)[:m]
		b := readString(reader)[:n]
		res := solve(a, b, k)

		buf.WriteString(fmt.Sprintf("%s\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a string, b string, k int) string {
	A := []byte(a)
	B := []byte(b)

	sort.Sort(Bytes(A))

	sort.Sort(Bytes(B))

	res := make([]Item, len(A)+len(B))
	var p int

	for i, j := 0, 0; i < len(A) && j < len(B); {
		if p == 0 {
			if A[i] < B[j] {
				res[p] = Item{i, 0, 0}
				i++
			} else {
				res[p] = Item{j, 0, 1}
				j++
			}
			p++
			continue
		}
		// p > 0
		if res[p-1].pos+1 == k {
			// have to choose another que
			if res[p-1].que == 0 {
				res[p] = Item{j, 0, 1}
				j++
			} else {
				res[p] = Item{i, 0, 0}
				i++
			}
			p++
			continue
		}

		if A[i] < B[j] {
			if res[p-1].que == 0 {
				res[p] = Item{i, res[p-1].pos + 1, 0}
			} else {
				res[p] = Item{i, 0, 0}
			}
			i++
		} else {
			if res[p-1].que == 1 {
				res[p] = Item{j, res[p-1].pos + 1, 1}
			} else {
				res[p] = Item{j, 0, 1}
			}
			j++
		}

		p++
	}

	buf := make([]byte, p)

	for i := 0; i < p; i++ {
		if res[i].que == 0 {
			buf[i] = A[res[i].idx]
		} else {
			buf[i] = B[res[i].idx]
		}
	}

	return string(buf)
}

type Item struct {
	idx int
	pos int
	que int
}

type Bytes []byte

func (this Bytes) Len() int {
	return len(this)
}

func (this Bytes) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Bytes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
