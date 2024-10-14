package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)

	fmt.Println(res)
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

const inf = 1 << 60

func solve(a []int) int {
	// 如果到i为止的sum 是k的倍数，那么它们没有道理继续往后移动
	// 所以，需要移动的只有刚好超过的那部分
	// 假设经过i以后，剩余了w个，且它们移动的距离为v
	// 那么, 那么下一个位置，也是确定的，就是正好 sum >= k 的地方
	// 但是这里的移动怎么计算呢？

	n := len(a)

	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + a[i]
	}

	sum := pref[n]
	if sum <= 1 {
		return -1
	}

	get := func(k int) int {
		var res int
		for i := 1; i < n; i++ {
			res += min(pref[i]%k, k-pref[i]%k)
		}
		return res
	}

	res := inf

	for k := 2; k <= sum/k; k++ {
		if sum%k == 0 {
			res = min(res, get(k))
			for sum%k == 0 {
				sum /= k
			}
		}
	}
	if sum > 1 {
		res = min(res, get(sum))
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
