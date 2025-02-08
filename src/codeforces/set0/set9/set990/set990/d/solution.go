package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, a, b := readThreeNums(reader)
	res := solve(n, a, b)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, a int, b int) []string {
	if a != 1 && b != 1 {
		return nil
	}
	if n == 1 {
		return []string{"0"}
	}
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			res[i][j] = '0'
		}
	}

	if b > 1 {
		// a == 1
		ans := solve(n, b, a)
		for i := range n {
			for j := i + 1; j < n; j++ {
				if ans[i][j] == '0' {
					res[i][j] = '1'
				} else {
					res[i][j] = '0'
				}
				res[j][i] = res[i][j]
			}
		}
		for i := range n {
			ans[i] = string(res[i])
		}
		return ans
	}

	if a+b == 2 {
		// 两个都为1，貌似是不行的
		if n <= 3 {
			return nil
		}
		// x (1, 3, 4) (2, 3)
		// (2, 1, 3) (1, 4)
		for j := 2; j < n; j++ {
			res[0][j] = '1'
			res[j][0] = '1'
		}
		res[1][2] = '1'
		res[2][1] = '1'
	} else {
		// b == 1
		// a > 1
		for j := a; j < n; j++ {
			res[a-1][j] = '1'
			res[j][a-1] = '1'
		}
	}
	ans := make([]string, n)
	for i := range n {
		ans[i] = string(res[i])
	}
	return ans
}
