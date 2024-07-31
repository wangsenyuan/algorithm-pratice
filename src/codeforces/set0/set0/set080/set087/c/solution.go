package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)
	fmt.Println(res)
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

func solve(n int) int {
	type pair struct{ a0, len int }
	ps := make([][]pair, n+1)
	sgSum := make([]int, n+2)

	for cnt := 2; cnt*(cnt+1)/2 <= n; cnt++ {
		for a := 1; ; a++ {
			s := (a + a + cnt - 1) * cnt / 2
			if s > n {
				break
			}
			ps[s] = append(ps[s], pair{a, cnt})
		}
	}

	for i := 3; i <= n; i++ {
		mex := make(map[int]bool)
		for _, p := range ps[i] {
			s := sgSum[p.a0] ^ sgSum[p.a0+p.len]
			if i == n && s == 0 {
				return p.len
			}
			mex[s] = true
		}
		var sg int
		for mex[sg] {
			sg++
		}
		sgSum[i+1] = sgSum[i] ^ sg
	}

	return -1
}
