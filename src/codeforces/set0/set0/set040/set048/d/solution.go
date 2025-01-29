package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	m := readNum(reader)
	a := readNNums(reader, m)
	n, res := solve(a)
	if n < 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", n))
	for i := range m {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	buf.WriteTo(os.Stdout)
}

func solve(a []int) (n int, res []int) {
	m := slices.Max(a)
	cnt := make([]int, m+1)
	for _, x := range a {
		cnt[x]++
	}
	n = cnt[1]
	res = make([]int, len(a))
	// cnt[1] >= cnt[2] >= cnt[3] ...只要满足这个条件就可以了
	for i := 2; i <= m; i++ {
		if cnt[i] > cnt[i-1] {
			return -1, nil
		}
	}
	for i := 0; i < len(a); i++ {
		x := a[i]
		res[i] = cnt[x]
		cnt[x]--
	}

	return
}
