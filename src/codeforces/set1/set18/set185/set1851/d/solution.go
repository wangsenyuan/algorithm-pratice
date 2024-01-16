package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		sums := readNNums(reader, n-1)
		res := solve(sums)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

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

func solve(prefs []int) bool {
	n := len(prefs) + 1

	sort.Ints(prefs)

	nums := make([]int, n+1)
	var special []int

	handleDiff := func(diff int) {
		if diff <= n {
			if nums[diff] == 0 {
				nums[diff]++
			} else {
				special = append(special, diff)
			}
		} else {
			special = append(special, diff)
		}
	}
	// 正常情况下，前缀和的diff，分布在1,2,3,4...n
	// 但是因为丢掉了一个，所以会出现某个数出现两次或者diff > n

	prev := 0
	for i := 0; i < len(prefs); i++ {
		diff := prefs[i] - prev
		if diff == 0 {
			return false
		}
		handleDiff(diff)
		prev = prefs[i]
	}
	last := n * (n + 1) / 2

	if prev > last {
		return false
	}

	if prev < last {
		handleDiff(last - prev)
	}
	if len(special) == 0 {
		return true
	}
	// 1, 2, 3, 4, 但是
	if len(special) > 1 {
		// 不能发生这种情况
		return false
	}
	diff := special[0]
	if diff >= 2*n {
		return false
	}

	// diff < 2 * n， 最大是 n + n - 1 (那么 nums[n] 和 nums[n-1]) 不应该出现
	var arr []int

	for i := 1; i <= n; i++ {
		if nums[i] == 0 {
			arr = append(arr, i)
		}
	}
	if len(arr) != 2 || arr[0]+arr[1] != special[0] {
		return false
	}

	return true
}
