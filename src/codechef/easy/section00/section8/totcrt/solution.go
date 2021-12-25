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
		n := readNum(reader)
		A := make([]string, n)
		for i := 0; i < n; i++ {
			A[i], _ = reader.ReadString('\n')
		}
		B := make([]string, n)
		for i := 0; i < n; i++ {
			B[i], _ = reader.ReadString('\n')
		}

		C := make([]string, n)
		for i := 0; i < n; i++ {
			C[i], _ = reader.ReadString('\n')
		}

		res := solve(n, A, B, C)

		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}

		buf.WriteByte('\n')
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

func solve(n int, A []string, B []string, C []string) []int64 {
	mem := make(map[string]int64)

	for i := 0; i < n; i++ {
		s, c := parse(A[i])
		mem[s] += int64(c)
		s, c = parse(B[i])
		mem[s] += int64(c)
		s, c = parse(C[i])
		mem[s] += int64(c)
	}
	arr := make([]int64, len(mem))
	var j int
	for _, v := range mem {
		arr[j] = v
		j++
	}
	sort.Sort(Int64s(arr))
	return arr
}

type Int64s []int64

func (this Int64s) Len() int {
	return len(this)
}

func (this Int64s) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64s) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func parse(s string) (string, int) {
	var i int
	for i < len(s) && s[i] != ' ' {
		i++
	}
	code := s[:i]
	var cnt int
	i++
	for i < len(s) {
		if s[i] == '\n' {
			break
		}
		cnt = cnt*10 + int(s[i]-'0')
		i++
	}
	return code, cnt
}
