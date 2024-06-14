package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(s string, k int) int {
	n := len(s)
	cnt := make([]int, 2)
	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		cnt[x]++
	}

	if cnt[0]%k != 0 {
		return -1
	}

	u := n / k

	if u%2 == 0 && cnt[0] != cnt[1] {
		return -1
	}
	if u%2 == 1 && abs(cnt[0]-cnt[1]) != k {
		return -1
	}

	var w int
	for i := n - 1; i >= 0; i-- {
		if s[i] == s[n-1] {
			w++
		} else {
			break
		}
	}

	if w > k {
		return -1
	}

	check := func(s string) bool {
		for i := 0; i < k; i++ {
			if s[i] != s[0] {
				return false
			}
		}
		for i := 0; i+k < n; i++ {
			if s[i] == s[i+k] {
				return false
			}
		}
		return true
	}

	if w == k {
		p := n
		for i := n - 1 - k; i >= 0; i-- {
			if s[i] == s[i+k] {
				p = i + 1
				break
			}
		}
		x := operate(s, p)
		if check(x) {
			return p
		}
		return -1
	}
	// w < k
	for i := 0; i < n; {
		if s[i] != s[n-1] {
			i++
			continue
		}
		j := i
		for i < n && s[i] == s[j] {
			i++
		}
		if i-j+w == k {
			x := operate(s, i)
			if check(x) {
				return i
			}
			return -1
		}
		if i-j+w == 2*k {
			x := operate(s, j+k-w)
			if check(x) {
				return j + k - w
			}
			return -1
		}
	}

	x := operate(s, n)
	if check(x) {
		return n
	}
	return -1
}

func operate(s string, p int) string {
	buf := []byte(s)
	reverse(buf[:p])
	tmp := string(buf[:p])
	copy(buf, buf[p:])
	copy(buf[len(s)-p:], tmp)
	return string(buf)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func abs(num int) int {
	return max(num, -num)
}
