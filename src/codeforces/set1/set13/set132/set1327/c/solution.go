package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	pos := make([][]int, k)
	for i := 0; i < k; i++ {
		pos[i] = readNNums(reader, 2)
	}
	target := make([][]int, k)
	for i := 0; i < k; i++ {
		target[i] = readNNums(reader, 2)
	}
	res := solve(n, m, pos, target)
	fmt.Println(len(res))
	fmt.Println(res)
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

const inf = 1 << 30

func solve(n int, m int, pos [][]int, target [][]int) string {

	res := strings.Repeat("U", n-1)
	res += strings.Repeat("L", m-1)

	for i := 0; i < n; i++ {
		if i > 0 {
			res += "D"
		}
		if i%2 == 0 {
			res += strings.Repeat("R", m-1)
		} else {
			res += strings.Repeat("L", m-1)
		}
	}
	return res
}
