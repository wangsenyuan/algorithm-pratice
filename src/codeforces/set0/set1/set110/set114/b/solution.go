package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	friends := make([]string, n)
	for i := range n {
		friends[i] = readString(reader)
	}
	dislikes := make([]string, m)
	for i := range m {
		dislikes[i] = readString(reader)
	}
	res := solve(friends, dislikes)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(friends []string, dislikes []string) []string {
	n := len(friends)
	id := make(map[string]int)
	for i, x := range friends {
		id[x] = i
	}

	var res int
	for state := 1; state < 1<<n; state++ {
		ok := true
		for _, cur := range dislikes {
			sp := strings.IndexByte(cur, ' ')
			x := cur[:sp]
			y := cur[sp+1:]
			i := id[x]
			j := id[y]
			if (state>>i)&1 > 0 && (state>>j)&1 > 0 {
				ok = false
				break
			}
		}
		if ok {
			if bits.OnesCount(uint(state)) > bits.OnesCount(uint(res)) {
				res = state
			}
		}
	}

	var ans []string

	for i := 0; i < n; i++ {
		if (res>>i)&1 == 1 {
			ans = append(ans, friends[i])
		}
	}
	sort.Strings(ans)
	
	return ans
}
