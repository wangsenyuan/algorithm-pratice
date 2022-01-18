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
		n, k := readTwoNums(reader)
		res := solve(n, k)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
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

func solve(n int, k int) []int {

	next := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			next[i/2+1] = i
		} else {
			next[i/2+n/2] = i
		}
	}

	ans := make([]int, n)

	vis := make([]bool, n)

	for i := 1; i <= n; i++ {
		if vis[i-1] {
			continue
		}
		var cycle []int
		cur := i
		vis[cur-1] = true
		cycle = append(cycle, cur)
		for next[cur] != i {
			cur = next[cur]
			vis[cur-1] = true
			cycle = append(cycle, cur)
		}

		cl := len(cycle)

		for j := 0; j < cl; j++ {
			ans[cycle[j]-1] = cycle[(j+k)%cl]
		}
	}

	return ans
}
