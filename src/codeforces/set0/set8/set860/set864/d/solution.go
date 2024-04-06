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
	nums := readNNums(reader, n)
	cnt, res := solve(nums)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", cnt))
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

func solve(a []int) (int, []int) {
	n := len(a)
	cnt := make([]int, n+1)
	var p int
	for i := 0; i < n; i++ {
		if a[i] > n || cnt[a[i]] >= 1 {
			p++
		}
		if a[i] <= n {
			cnt[a[i]]++
		}
	}

	// res怎么计算呢
	res := make([]int, n)
	expect := 1
	for i := 0; i < n; i++ {
		if a[i] <= n && cnt[a[i]] == 1 {
			res[i] = a[i]
			continue
		}

		// 如果 a[i] 是bad的
		for expect <= n && cnt[expect] >= 1 {
			expect++
		}

		if a[i] > n || expect < a[i] || cnt[a[i]] == 0 {
			// a[i]之前已经被处理了
			res[i] = expect
			if a[i] <= n && cnt[a[i]] > 0 {
				cnt[a[i]]--
			}
			expect++
		} else {
			res[i] = a[i]
			cnt[a[i]] = 0
		}
	}

	return p, res
}

type Pair struct {
	first  int
	second int
}
