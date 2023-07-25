package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line := readNNums(reader, 4)

	res := solve(line[0], line[1], line[2], line[3])

	if len(res) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(res)
	}
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

func solve(n int, k int, a int, b int) string {
	if (a == 0 || b == 0) && k < n {
		return ""
	}
	if a < b {
		res := solve(n, k, b, a)
		if len(res) > 0 {
			res = flip(res)
		}
		return res
	}
	// a > b
	buf := make([]byte, n)

	// 然后使用 GB 进行配对， 一共是b对，还剩余 a - b个g
	diff := a - b
	// 此时有 b 个 空间， 每个里面最多放置k-1 G, 然后是一个 GB, 最后一个可以放置k个G
	var it int
	for i := 0; i < b; i++ {
		for j := 0; j < k-1 && diff > 0 && it < n; j++ {
			buf[it] = 'G'
			it++
			diff--
		}
		// 至少要吧b个B放置掉
		if it == n {
			return ""
		}
		buf[it] = 'G'
		it++
		if it == n {
			return ""
		}
		buf[it] = 'B'
		it++
	}
	if diff > k || diff != n-it {
		return ""
	}

	for it < n {
		buf[it] = 'G'
		it++
	}

	return string(buf)
}

func flip(s string) string {
	buf := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'B' {
			buf[i] = 'G'
		} else {
			buf[i] = 'B'
		}
	}
	return string(buf)
}
