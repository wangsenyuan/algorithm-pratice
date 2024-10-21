package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	cnt := readNNums(reader, 4)

	res := solve(cnt)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	s := fmt.Sprintf("%v", res)

	fmt.Println("YES")
	fmt.Println(s[1 : len(s)-1])
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

func solve(cnt []int) []int {

	var n int
	for i := 0; i < 4; i++ {
		n += cnt[i]
	}

	tmp := make([]int, 4)

	process := func(cur int) []int {
		if cnt[cur] == 0 {
			return nil
		}
		copy(tmp, cnt)
		var res []int
		res = append(res, cur)
		tmp[cur]--

		for len(res) < n {
			if cur > 0 && tmp[cur-1] > 0 {
				cur--
				tmp[cur]--
				res = append(res, cur)
			} else if cur+1 < 4 && tmp[cur+1] > 0 {
				cur++
				tmp[cur]--
				res = append(res, cur)
			} else {
				break
			}
		}
		if len(res) == n {
			return res
		}
		return nil
	}

	for i := 0; i < 4; i++ {
		res := process(i)
		if len(res) > 0 {
			return res
		}
	}

	return nil
}
