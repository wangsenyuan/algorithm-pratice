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
		n, m, q := readThreeNums(reader)
		E := make([]int, q)
		for i := 0; i < q; i++ {
			s, _ := reader.ReadBytes('\n')
			var j int
			readInt(s, 2, &j)
			if s[0] == '+' {
				E[i] = j
			} else {
				E[i] = -j
			}
		}
		res := solve(n, m, E)

		if res {
			buf.WriteString("Consistent\n")
		} else {
			buf.WriteString("Inconsistent\n")
		}
	}

	fmt.Print(buf.String())
}

func solve(N, M int, E []int) bool {
	in := make([]bool, N+1)
	var cnt int
	for _, cur := range E {
		if cur > 0 {
			in[cur] = true
			cnt++
		} else {
			cur = -cur
			if !in[cur] {
				return false
			}
			in[cur] = false
			cnt--
		}
		if cnt > M {
			return false
		}
	}
	return true
}
