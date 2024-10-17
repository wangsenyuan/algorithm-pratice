package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	src := make([]string, n)
	for i := 0; i < n; i++ {
		src[i] = readString(reader)
	}
	tgt := make([]string, n)
	for i := 0; i < n; i++ {
		tgt[i] = readString(reader)
	}
	bar := readString(reader)
	ok, res := solve(src, tgt, bar)
	if !ok {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}

	fmt.Print(res)
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

func solve(src []string, tgt []string, bar string) (bool, []string) {
	n := len(src)

	g := make([][]int32, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int32, n)
		for j := 0; j < n; j++ {
			if src[i][j] != tgt[i][j] {
				g[i][j] = 1
			}
		}
	}

	id := -1
	b := make([]int32, n)
	for i := 0; i < n; i++ {
		b[i] = int32(bar[i] - '0')
		if b[i] == 1 && id < 0 {
			id = i
		}
	}
	if id < 0 {
		return false, nil
	}

	var ans []string

	for i := 0; i < n; i++ {
		if g[id][i] == 1 {
			ans = append(ans, fmt.Sprintf("col %d", i))
			for j := 0; j < n; j++ {
				g[j][i] ^= b[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		if g[i][id] == 1 {
			ans = append(ans, fmt.Sprintf("row %d", i))
			for j := 0; j < n; j++ {
				g[i][j] ^= b[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] == 1 {
				return false, nil
			}
		}
	}

	return true, ans
}
