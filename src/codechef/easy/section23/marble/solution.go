package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		S, _ := reader.ReadString('\n')
		P, _ := reader.ReadString('\n')
		res := solve(n, S, P)
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

func solve(n int, S, P string) int {

	get := func(x, y byte) byte {
		if x == '?' {
			return y
		}
		return x
	}

	count := func(x byte) int {
		var res int
		for i := 0; i < n; i++ {
			a := get(S[i], x)
			b := get(P[i], x)
			if a == b {
				continue
			}
			// a != b
			if isVowel(a) != isVowel(b) {
				res++
			} else {
				res += 2
			}
		}
		return res
	}
	best := n * 2
	for x := 'a'; x <= 'z'; x++ {
		best = min(best, count(byte(x)))
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func isVowel(x byte) bool {
	return x == 'a' || x == 'o' || x == 'e' || x == 'i' || x == 'u'
}
