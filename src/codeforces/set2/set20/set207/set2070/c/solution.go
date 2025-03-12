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
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	s := readString(reader)
	a := readNNums(reader, n)
	return solve(k, s, a)
}

func solve(k int, s string, a []int) int {
	b := sortAndUnique(a)
	n := len(a)

	check := func(x int) bool {
		// 只需要处理连续的蓝色(超过x)的段
		var cnt int
		for i := 0; i < n; {
			if a[i] < x || s[i] == 'R' {
				i++
				continue
			}
			// a[i] >= k and a[i] == 'B'
			// 必须开始一段新的
			cnt++
			i++
			for i < n && (s[i] == 'B' || a[i] < x) {
				i++
			}
			// i == n or a[i] >= x && s[i] == 'R', 结束一段
		}

		return cnt <= k
	}

	i := sort.Search(len(b), func(i int) bool {
		return check(b[i])
	})
	if i == 0 {
		return 0
	}
	return b[i-1]
}

func sortAndUnique(a []int) []int {
	arr := make([]int, len(a))
	copy(arr, a)
	sort.Ints(arr)
	var it int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i-1] < arr[i] {
			arr[it] = arr[i-1]
			it++
		}
	}
	return arr[:it]
}
