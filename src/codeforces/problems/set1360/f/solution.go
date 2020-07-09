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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		S := make([][]byte, n)
		for i := 0; i < n; i++ {
			S[i], _ = reader.ReadBytes('\n')
		}
		res := solve(n, m, S)
		if len(res) == 0 {
			fmt.Println("-1")
		} else {
			fmt.Println(res)
		}
	}
}

func solve(n, m int, A [][]byte) string {
	buf := make([]byte, m)

	var dfs func(flag int, i int) bool

	dfs = func(flag int, i int) bool {
		if i == m {
			return true
		}

		set := make(map[byte]bool)

		for j := 0; j < n; j++ {
			set[A[j][i]] = true
		}

		for k := range set {
			var flag2 int

			for j := 0; j < n; j++ {
				if A[j][i] != k {
					flag2 |= 1 << j
				}
			}
			if flag2&flag > 0 {
				continue
			}
			buf[i] = k
			if dfs(flag|flag2, i+1) {
				return true
			}
		}
		return false
	}

	res := dfs(0, 0)

	if res {
		return string(buf)
	}

	return ""
}
