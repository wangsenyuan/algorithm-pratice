package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	res := solve(n, k)

	if len(res) == 0 {
		fmt.Println("No")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("Yes\n")
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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
	var cnt int
	res := make([]int, 151)
	for i := 62; i >= 0; i-- {
		if (n>>i)&1 == 1 {
			res[i+62]++
			cnt++
		}
	}
	if cnt > k {
		return nil
	}

	count := func(x int) int {
		var tot int
		for i := 150; i > x; i-- {
			c := res[i]
			tot += c * (1 << (i - x))
		}

		for i := x; i >= 0; i-- {
			tot += res[i]
		}

		return tot
	}

	l, r := 0, 150

	for l < r {
		mid := (l + r) >> 1
		if count(mid) <= k {
			r = mid
		} else {
			l = mid + 1
		}
	}

	for i := 150; i > r; i-- {
		res[r] += res[i] * (1 << (i - r))
	}
	// 应该减少小的那些，而不是大的
	var ans []int

	for i := r; i >= 0; i-- {
		for res[i] > 0 {
			ans = append(ans, i-62)
			res[i]--
		}
	}

	for len(ans) < k {
		m := len(ans)
		x := ans[m-1]
		ans[m-1] = x - 1
		ans = append(ans, x-1)
	}

	return ans
}
