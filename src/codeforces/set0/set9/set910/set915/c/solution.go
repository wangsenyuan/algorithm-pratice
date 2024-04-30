package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := readNum(reader)
	b := readNum(reader)

	res := solve(a, b)

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

func solve(a, b int) int {
	if a == b {
		return a
	}

	as := getDigits(a)
	bs := getDigits(b)

	if len(as) < len(bs) {
		return solve1(as)
	}

	n := len(bs)

	cnt := make([]int, 10)
	for _, x := range as {
		cnt[x]++
	}

	tmp := make([]int, 10)
	// 先要确定， 前面多少位, a = b
	// 必须在满足条件的情况下， a <= b

	check := func(d int) bool {
		copy(tmp, cnt)
		// 后面d位 a[i] = b[i]
		for i := 0; i < d; i++ {
			j := n - 1 - i
			x := bs[j]
			if tmp[x] == 0 {
				return false
			}
			tmp[x]--
		}
		// 只要保证前d位相同，且能够使得 a <= b 即可
		for i := d; i < n; i++ {
			x := bs[n-1-i]
			for j := 0; j < x; j++ {
				if tmp[j] > 0 {
					return true
				}
			}
			if tmp[x] == 0 {
				return false
			}
			tmp[x]--
		}

		return true
	}

	if check(n) {
		// can get b
		return b
	}

	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// check(r) = false => r太大了
	r--
	// check(r) = true
	var res int
	for i := 0; i < r; i++ {
		res = res*10 + bs[n-1-i]
		cnt[bs[n-1-i]]--
	}
	x := bs[n-1-r] - 1
	for cnt[x] == 0 {
		x--
	}
	cnt[x]--
	res = res*10 + x
	for i := 9; i >= 0; i-- {
		for cnt[i] > 0 {
			res = res*10 + i
			cnt[i]--
		}
	}

	return res
}

func solve1(arr []int) int {
	sort.Ints(arr)
	var res int
	for i := len(arr) - 1; i >= 0; i-- {
		res = res*10 + arr[i]
	}
	return res
}

func getDigits(num int) []int {
	var res []int
	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}

	return res
}
