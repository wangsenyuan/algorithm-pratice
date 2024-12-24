package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		res := solve(n, k)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d ", x))
			}
			buf.WriteByte('\n')
		}
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

func solve(n int, k int) []int {
	if n <= 60 && 1<<(n-1) < k {
		return nil
	}
	// else must have anser
	res := make([]int, n)
	// 1,2,3... 必须在两边
	l, r := 0, n-1
	for i := 1; i <= n; i++ {
		// 如果i放在l位置上，那么必须要求 k <= 1 << (r - l)
		if (r-l) > 60 || l == r || (1<<(r-l-1)) >= k {
			res[l] = i
			l++
		} else {
			k -= 1 << (r - l - 1)
			res[r] = i
			r--
		}
	}
	return res
}
