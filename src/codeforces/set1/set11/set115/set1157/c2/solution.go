package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	//tc := readNum(scanner)
	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)

		buf.WriteString(fmt.Sprintf("%d\n%s\n", len(res), res))
	}

	fmt.Println(buf.String())
}

func solve(que []int) string {
	// a  b, c
	// 假设 a < c, 是不是取a一定是更优的方案呢？貌似是的
	// 但是如果 a == c, 那这时后就有问题了
	// 如果取走a，c肯定是不能被取走了，对c也是一样的
	// 但是只需要检查a后面和c后面的长度即可
	var buf bytes.Buffer

	n := len(que)
	l, r := 0, n-1
	prev := -1
	for l < r && max(que[l], que[r]) > prev {
		if que[r] <= prev || que[l] < que[r] && que[l] > prev {
			buf.WriteByte('L')
			prev = que[l]
			l++
		} else if que[l] <= prev || que[l] > que[r] && que[r] > prev {
			buf.WriteByte('R')
			prev = que[r]
			r--
		} else {
			break
		}
	}
	if l == r {
		if que[l] > prev {
			buf.WriteByte('L')
		}
		return buf.String()
	}

	// l < r  que[l] == que[r]
	cnt := make([]int, 2)
	if que[l] > prev {
		cnt[0]++
		for i := l; i < r && que[i+1] > que[i]; i++ {
			cnt[0]++
		}
	}
	if que[r] > prev {
		cnt[1]++
		for i := r; i > l && que[i-1] > que[i]; i-- {
			cnt[1]++
		}
	}

	if cnt[0] >= cnt[1] {
		for i := 0; i < cnt[0]; i++ {
			buf.WriteByte('L')
		}
	} else {
		for i := 0; i < cnt[1]; i++ {
			buf.WriteByte('R')
		}
	}

	return buf.String()
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
