package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ask := func(t int, i int, j int, x int) int {
		fmt.Printf("? %d %d %d %d\n", t, i, j, x)
		return readNum(reader)
	}

	tc := readNum(reader)
	for range tc {
		n := readNum(reader)
		res := solve(n, ask)
		var buf bytes.Buffer
		buf.WriteString("!")
		for _, x := range res {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
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

func solve(n int, ask func(t int, i int, j int, x int) int) []int {
	var one int
	p := make([]int, n+1)
	for i := 1; i <= n; i += 2 {
		j := i + 1
		if j > n {
			j = 1
		}
		ans := ask(2, i, j, 1)
		if ans == 1 {
			one = i
			break
		}
		if ans == 2 {
			// min(max(1, pi), max(2, pj)) = 2
			// 如果p[j] = 1, pi >= 2
			// 或者p[i] = 2, pj >= 2
			tmp := ask(2, j, i, 1)
			if tmp == 1 {
				// 这个肯定能找到
				one = j
				break
			}
			// 说明p[j] >= 2
			// p[i] = 2 肯定成立吗？
		}
	}

	p[one] = 1

	for i := 1; i <= n; i++ {
		if i == one {
			continue
		}
		// max(min(n-1, 1), min(n, pj)) = p[j]
		p[i] = ask(1, one, i, n-1)
	}

	return p[1:]
}
