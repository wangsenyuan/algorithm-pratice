package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)

	res := solve(a, k)

	fmt.Println(res)
}

func solve(a []int, k int) int {
	n := len(a)
	sort.Ints(a)

	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = a[i] + suf[i+1]
	}

	pref := make([]int, n+1)
	for i := range a {
		pref[i+1] = pref[i] + a[i]
	}

	check := func(expect int) bool {
		var sum int
		// 如果i是下边界
		for i, r := 0, 0; i < n; i++ {
			for r < n && a[r]-a[i] <= expect {
				r++
			}
			// 后面的都增加到a[i]
			cnt := i*a[i] - sum
			cnt += suf[r] - (n-r)*(a[i]+expect)
			if cnt <= k {
				return true
			}
			sum += a[i]
		}
		// 如果i是上边界
		sum = 0
		for i, l := n-1, n-1; i >= 0; i-- {
			for l >= 0 && a[i]-a[l] <= expect {
				l--
			}
			cnt := sum - (n-i-1)*a[i]
			cnt += (l+1)*(a[i]-expect) - pref[l+1]
			if cnt <= k {
				return true
			}

			sum += a[i]
		}

		return false
	}

	worst := a[n-1] - a[0]

	return sort.Search(worst, check)
}
