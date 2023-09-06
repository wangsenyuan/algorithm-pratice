package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)
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

const H = 30

func solve(a, b []int) int {
	//n := len(a)
	// try bit from high to low
	// a[i] ^ b[i] & ...
	// 对于数组x,y 存在一个匹配，使得d位为1
	// 对于d位，如何保证它一定可以为1呢
	// 对于最高位，确实可以按照d分为两组，0和1，然后如果数量一致，就可以在得到1
	// 那么对于下一位呢？需要保证高位不影响的情况下，分成两组即可
	check := func(mask int) bool {
		cnt := make(map[int]int)
		for _, x := range a {
			cnt[x&mask]++
		}
		for _, x := range b {
			cnt[(^x)&mask]--
		}
		for _, v := range cnt {
			if v != 0 {
				return false
			}
		}
		return true
	}

	var ans int
	for i := H - 1; i >= 0; i-- {
		if check(ans | (1 << i)) {
			ans |= (1 << i)
		}
	}
	return ans
}
