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
	var sum int
	for _, num := range a {
		sum += num
	}
	if sum <= 1 {
		return -1
	}

	n := len(a)

	get := func(k int) int {
		var res int
		for i := 0; i < n; i++ {
			if a[i] == 0 {
				continue
			}
			j := i
			var tmp int
			var suf int
			for i < n && tmp < k {
				if a[i] == 1 {
					// 这个移动到位置j处
					suf += i
					tmp++
				}
				i++
			}
			// tmp == k
			// 需要把[j...i)的都移动到一个位置上
			// 移动到中点
			best := suf - k*j
			var pref int
			var cnt int
			for j < i {
				// 所有的都移动到位置j处
				if a[j] == 1 {
					tmp--
					suf -= j
					pref += j
					cnt++
				}
				best = min(best, cnt*j-pref+suf-tmp*j)
				j++
			}
			res += best
			i--
		}
		return res
	}

	res := inf

	for k := 1; k <= sum/k; k++ {
		if sum%k == 0 {
			if k > 1 {
				res = min(res, get(k))
			}
			if k*k != sum && sum/k > 1 {
				res = min(res, get(sum/k))
			}
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
