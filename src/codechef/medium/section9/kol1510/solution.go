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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		K := readNum(reader)
		A := make([]string, K)
		for i := 0; i < K; i++ {
			A[i], _ = reader.ReadString('\n')
			A[i] = A[i][:len(A[i])-1]
		}
		fmt.Println(solve(K, A))
	}
}

func solve(K int, necklaces []string) string {

	var x int

	for i := 0; i < K; i++ {
		if x < len(necklaces[i]) {
			x = len(necklaces[i])
		}
	}

	buf := make([]byte, 16)

	check := func(n int) bool {
		for i := 0; i < K; i++ {
			if !isSubSeq(necklaces[i], buf[:n]) {
				return false
			}
		}
		return true
	}

	var dfs func(i int, n int) bool

	dfs = func(i int, n int) bool {
		if i == n {
			return check(n)
		}

		buf[i] = 'B'
		if dfs(i+1, n) {
			return true
		}
		buf[i] = 'G'

		return dfs(i+1, n)
	}

	find := func(n int) bool {
		return dfs(0, n)
	}

	for !find(x) {
		x++
	}

	return string(buf[:x])
}

func isSubSeq(s string, buf []byte) bool {
	if len(s) > len(buf) {
		return false
	}

	var i, j int

	for i < len(s) && j < len(buf) {
		if s[i] == buf[j] {
			i++
		}
		j++
	}

	return i == len(s)
}
