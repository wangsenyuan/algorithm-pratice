package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = readString(reader)
	}
	res := solve(s)
	fmt.Println(res)
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

func solve(s []string) int64 {
	// B**, **A, B**A
	cnt := make([]int, 3)
	var ans int64
	for _, x := range s {
		ans += int64(check(x))
		if x[0] == 'B' {
			if x[len(x)-1] == 'A' {
				cnt[2]++
			} else {
				cnt[0]++
			}
		} else if x[len(x)-1] == 'A' {
			cnt[1]++
		}
	}

	// **A B**A
	// B**A 可以变成一个**A
	if cnt[2] > 0 {
		if cnt[1] > 0 {
			// match it with the first B...A
			cnt[1]--
			ans++
		}

		ans += int64(cnt[2] - 1)
		cnt[1]++
	}

	ans += int64(min(cnt[1], cnt[0]))

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func check(s string) int {
	var ans int
	for i := 0; i+1 < len(s); i++ {
		if s[i] == 'A' && s[i+1] == 'B' {
			ans++
		}
	}
	return ans
}
