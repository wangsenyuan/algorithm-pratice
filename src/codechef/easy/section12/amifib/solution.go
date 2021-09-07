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
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		var num uint64
		s, _ := reader.ReadBytes('\n')
		readUint64(s, 0, &num)
		res := solve(num)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

const D = 6666

var fp []uint64

func init() {
	fp = make([]uint64, 0, D)
	fp = append(fp, 0)
	fp = append(fp, 1)
	for i := 2; i < D; i++ {
		cur := fp[i-2] + fp[i-1]
		fp = append(fp, cur)
	}
	sort.Sort(Uint64s(fp))
}

type Uint64s []uint64

func (this Uint64s) Len() int {
	return len(this)
}

func (this Uint64s) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Uint64s) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func solve(n uint64) bool {
	i := sort.Search(len(fp), func(i int) bool {
		return fp[i] >= n
	})
	if i == D || fp[i] != n {
		return false
	}
	return true
}
