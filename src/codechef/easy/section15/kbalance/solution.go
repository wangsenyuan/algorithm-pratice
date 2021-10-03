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
		n, k := readTwoNums(reader)
		S, _ := reader.ReadString('\n')
		res := solve(n, k, S)
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

func solve(n int, k int, S string) int {
	// if s[i] == '1', and it violates the condition
	// then one way is to turn it off, else turn another bit on
	// s[i] == '1'，S[j] == '0', and S[k] == '0'
	bad := make([]bool, n)

	for i := 0; i < n; i++ {
		if S[i] == '0' {
			bad[i] = false
			continue
		}
		if i-k >= 0 && S[i-k] == '1' {
			bad[i] = false
		} else if i+k < n && S[i+k] == '1' {
			bad[i] = false
		} else {
			bad[i] = true
		}
	}

	var res int

	for i := 0; i < n; i++ {
		if bad[i] {
			// a one
			res++
			if i+2*k < n {
				bad[i+2*k] = false
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
