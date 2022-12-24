package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if len(res) == 0 {
			buf.WriteString("-1")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
		}

		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

var primes []int
var lpf []int

const MAX_A = 1000010

func init() {

	lpf = make([]int, MAX_A)

	for i := 2; i < MAX_A; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			for j := i; j < MAX_A; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func solve(B []int) []int {
	n := len(B)
	//
	// B[0] = x, A[0]
	pos := make(map[int][]int)

	id := make(map[int]int)

	addFreq := func(f int, i int) {
		if pos[f] == nil {
			pos[f] = make([]int, 0, 1)
		}
		pos[f] = append(pos[f], i)
	}

	for i, f := range B {
		addFreq(f, i)
	}

	A := make([]int, n)

	num := 1
	for i := 0; i < n; i++ {
		if A[i] > 0 {
			continue
		}
		f := B[i]
		cnt := f
		v := pos[f]

		for id[f] < len(v) && cnt > 0 {
			A[v[id[f]]] = num
			id[f]++
			cnt--
		}
		if cnt > 0 {
			return nil
		}
		num++
	}
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			return nil
		}
	}
	return A
}
