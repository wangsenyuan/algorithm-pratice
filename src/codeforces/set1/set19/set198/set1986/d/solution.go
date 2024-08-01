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
		readNum(reader)
		s := readString(reader)
		res := solve(s)
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

func solve(s string) int {
	n := len(s)

	if n == 2 {
		return int(s[0]-'0')*10 + int(s[1]-'0')
	}

	best := 1 << 30

	// c0 == 0 || c0 == 1 and n == 3
	arr := make([]int, n-1)
	for i := 0; i+1 < n; i++ {
		// 如果在ii+1处连接在一起
		for j, k := 0, 0; j < n; j++ {
			if i == j {
				arr[k] = int(s[j]-'0')*10 + int(s[j+1]-'0')
				j++
			} else {
				arr[k] = int(s[j] - '0')
			}
			k++
		}
		var sum int
		one := true
		for j := 0; j < n-1; j++ {
			if arr[j] == 0 {
				return 0
			}
			if arr[j] != 1 {
				sum += arr[j]
			}
			one = one && arr[j] == 1
		}
		if one {
			best = 1
		} else {
			best = min(best, sum)
		}
	}

	return best
}

func count(s string, x byte) int {
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == x {
			res++
		}
	}
	return res
}
