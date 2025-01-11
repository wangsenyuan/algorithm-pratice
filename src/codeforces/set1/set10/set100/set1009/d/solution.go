package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	res := solve(n, m)
	if len(res) == 0 {
		fmt.Println("Impossible")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("Possible\n")
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

func solve(n int, m int) [][]int {
	if m < n-1 {
		return nil
	}
	phi := make([]int, n+1)
	phi[1] = 1
	for i := 2; i <= n; i++ {
		phi[i] = i - 1
	}
	for i := 2; i <= n; i++ {
		for j := i + i; j <= n; j += i {
			phi[j] -= phi[i]
		}
	}
	var sum int
	for i := 2; i <= n; i++ {
		sum += phi[i]
	}
	// 1能到1， 所以要减去1
	if sum < m {
		return nil
	}
	// phi[n] - 1 >= m
	// 考虑4和6

	var res [][]int

	for i := 1; i < n && len(res) < m; i++ {
		for j := i + 1; j <= n && len(res) < m; j++ {
			if gcd(i, j) == 1 {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
