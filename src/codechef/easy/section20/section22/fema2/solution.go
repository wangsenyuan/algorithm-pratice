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
		n, K := readTwoNums(reader)
		s, _ := reader.ReadString('\n')
		res := solve(n, K, s)
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

func solve(n int, K int, S string) int {
	var res int
	var j int
	stack := make([]int, n)
	sheet := make([]int, n)
	for i := 0; i <= n; i++ {
		if i == n || S[i] == 'X' {
			res += count(S[j:i], K, stack, sheet)
			j = i + 1
		}
	}

	return res
}

func count(s string, K int, stack []int, sheet []int) int {
	// K + 1 - |j - i| - sij
	var p int
	sheet[0] = 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			stack[p] = i
			p++
		}
		sheet[i] = 0
		if s[i] == ':' {
			sheet[i]++
		}
		if i > 0 {
			sheet[i] += sheet[i-1]
		}
	}
	var res int
	var j int
	for i := 0; i < len(s); i++ {
		if s[i] == 'M' {
			var ok bool
			for j < p && stack[j] < i {
				tmp := K + 1 - (i - stack[j]) - (sheet[i] - sheet[stack[j]])
				j++
				if tmp > 0 {
					ok = true
					res++
					break
				}
			}
			if !ok && j < p {
				// stack[j] > i
				tmp := K + 1 - (stack[j] - i) - (sheet[stack[j]] - sheet[i])
				if tmp > 0 {
					j++
					res++
				}
			}
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
