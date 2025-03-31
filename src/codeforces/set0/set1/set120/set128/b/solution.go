package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	k := readNum(reader)
	res := solve(s, k)
	fmt.Println(res)
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
	return strings.TrimSpace(s)
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func solve(s string, k int) string {
	n := len(s)
	tot := n * (n + 1) / 2
	if tot < k {
		return "No such line."
	}
	pos := make([][]int, K)
	cnt := make([]int, K)
	for i := range n {
		x := int(s[i] - 'a')
		cnt[x] += n - i
		pos[x] = append(pos[x], i)
	}

	var buf bytes.Buffer
	for k > 0 {
		var sum int
		var x int
		for sum+cnt[x] < k {
			sum += cnt[x]
			x++
		}
		// sum + cnt[x] >= k
		buf.WriteByte(byte(x + 'a'))
		k -= sum
		vs := pos[x]
		k -= len(vs)
		clear(pos)
		clear(cnt)
		for _, i := range vs {
			i++
			if i < n {
				u := int(s[i] - 'a')
				cnt[u] += n - i
				pos[u] = append(pos[u], i)
			}
		}
	}

	return buf.String()
}

const K = 26

func bruteForce(s string, k int) string {
	var x []string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			x = append(x, s[i:j+1])
		}
	}
	sort.Strings(x)

	return x[k-1]
}
