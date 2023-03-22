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

	for tc > 0 {
		tc--
		n := readNum(reader)
		nums := readNNums(reader, n)
		res := solve(nums)
		for _, num := range res {
			buf.WriteString(fmt.Sprintf("%d ", num))
		}
		buf.WriteByte('\n')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A []int) []int64 {
	n := len(A)
	// A[i] 只能增加，不能减少

	sort.Ints(A)

	var B []int

	for i := 1; i <= n; i++ {
		if i == n || A[i] > A[i-1] {
			B = append(B, A[i-1])
		}
	}

	// 如果要使得mex = x, 那么所有小于x的数字必须出现，且 = x的数字必须加1
	cnt := make([]int, n+1)
	for i := 0; i < n && A[i] <= n; i++ {
		cnt[A[i]]++
	}

	stack := make([]int, n+1)
	var p int

	var sum int64

	res := make([]int64, n+1)

	for i := 0; i <= n; i++ {
		res[i] = -1
	}

	for x, i := 0, 0; x <= n; x++ {
		// all need to add 1
		if x > 0 && cnt[x-1] == 0 {
			// to fill x - 1
			if p == 0 {
				// no nums to fill this gap
				break
			}
			sum += int64(x - 1 - stack[p-1])
			p--
		}

		res[x] = sum + int64(cnt[x])

		if i < len(B) && x == B[i] {
			for cnt[x] > 1 {
				cnt[x]--
				stack[p] = x
				p++
			}
			i++
		}
	}

	return res
}
