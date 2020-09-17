package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadBytes('\n')
	var k uint64
	readUint64(s, 0, &k)
	res := solve(int64(k))
	fmt.Println(res)
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

const codeforces = "codeforces"

func solve(k int64) string {
	n := len(codeforces)

	var c int = 1
	cnt := make([]int, n)
	for {
		var prod int64 = 1
		for i := 0; i < n; i++ {
			cnt[i] = c
			prod *= int64(c)
		}
		for i := n - 1; i >= 0; i-- {
			//chang i + 1 to c + 1
			if prod >= k {
				return createCodeforces(cnt)
			}
			cnt[i]++
			prod /= int64(c)
			prod *= int64(c + 1)
		}

		c++
	}
}

func product(cnt []int) int64 {
	var res int64 = 1
	for i := 0; i < len(cnt); i++ {
		res *= int64(cnt[i])
	}
	return res
}

func createCodeforces(cnt []int) string {
	var buf bytes.Buffer

	for i := 0; i < len(codeforces); i++ {
		for x := 0; x < cnt[i]; x++ {
			buf.WriteByte(codeforces[i])
		}
	}

	return buf.String()
}

func pow(a int64, b int) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}
