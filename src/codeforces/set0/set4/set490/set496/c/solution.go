package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	a := make([]string, n)

	for i := 0; i < n; i++ {
		a[i] = readString(reader)[:m]
	}

	res := solve(a)

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
func solve(a []string) int {
	n := len(a)
	m := len(a[0])

	var res int

	// ok[i] 表示，这一行是否已经比下一行小了
	ok := make([]bool, n)

	for j := 0; j < m; j++ {
		keep := true
		for i := 0; i+1 < n; i++ {
			if ok[i] {
				continue
			}
			if a[i][j] > a[i+1][j] {
				// 只要存在, 就必须删除
				keep = false
				break
			}
		}
		if !keep {
			res++
			continue
		}
		for i := 0; i+1 < n; i++ {
			if !ok[i] {
				ok[i] = a[i][j] < a[i+1][j]
			}
		}
	}

	return res
}
