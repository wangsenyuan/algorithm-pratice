package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s, _ := reader.ReadString('\n')
		res := solve(s[:len(s)-1])
		fmt.Printf("%d\n", res)
	}
}

func solve(s string) int {
	//fmt.Fprintln(os.Stderr, s)
	// all 0, all 1,
	// half 0, half 1
	var one int
	n := len(s)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		if s[i] != '0' && s[i] != '1' {
			n = i
			break
		}
		one += int(s[i] - '0')
		cnt[i] = one
	}

	ans := min(cnt[n-1], n-cnt[n-1])

	for i := n - 2; i >= 0; i-- {
		right := cnt[n-1] - cnt[i]
		left := cnt[i]
		// 01
		ans = min(ans, n-1-i-right+left)
		// or 10
		ans = min(ans, right+(i+1)-left)
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
