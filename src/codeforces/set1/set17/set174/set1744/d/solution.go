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
		a := readNNums(reader, n)
		res := solve(a)
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

const N = 200010

var f2 [N]int

func init() {
	f2[2] = 1
	for i := 3; i < N; i++ {
		if i%2 == 0 {
			f2[i] = 1 + f2[i/2]
		}
	}
}

func solve(a []int) int {
	n := len(a)
	// i * a[i] divisible by 2 ** n
	var cnt int

	arr := make([]int, n)

	for i, x := range a {
		for x&1 == 0 {
			cnt++
			x >>= 1
		}
		arr[i] = x
	}

	if cnt >= n {
		// already good
		return 0
	}

	tmp := make([]int, n+1)
	copy(tmp, f2[:n+1])
	sort.Ints(tmp[1:])

	var ans int

	for i := n; i > 0 && cnt < n; i-- {
		if tmp[i] == 0 {
			break
		}
		ans++
		cnt += tmp[i]
	}
	if cnt < n {
		return -1
	}
	return ans
}
