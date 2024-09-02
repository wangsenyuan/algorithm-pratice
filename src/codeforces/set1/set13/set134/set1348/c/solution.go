package main

import (
	"bufio"
	"bytes"
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, k)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func solve(s string, k int) string {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[int(s[i]-'a')]++
	}

	var arr []int
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			arr = append(arr, i)
		}
	}

	n := len(s)

	if len(arr) == 1 {
		return solve1(n, k, arr[0])
	}

	if len(arr) == 2 && freq[arr[0]] == k {
		w := solve1(n-k, k, arr[1])
		return fmt.Sprintf("%c%s", byte('a'+arr[0]), w)
	}

	if freq[arr[0]] < k {
		var x int
		var sum int
		for x < 26 && sum < k {
			sum += freq[x]
			x++
		}
		// sum >= k
		x--
		return fmt.Sprintf("%c", byte(x+'a'))
	}

	freq[arr[0]] -= k

	var buf bytes.Buffer
	buf.WriteByte(byte(arr[0] + 'a'))

	for _, x := range arr {
		for freq[x] > 0 {
			buf.WriteByte(byte(x + 'a'))
			freq[x]--
		}
	}

	return buf.String()
}

func solve1(n int, k int, x int) string {
	m := (n + k - 1) / k
	buf := make([]byte, m)
	for i := 0; i < m; i++ {
		buf[i] = byte(x + 'a')
	}
	return string(buf)
}
