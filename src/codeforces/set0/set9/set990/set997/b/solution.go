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

func solve(n int) int {
	m := min(n, 13)
	set := make(map[int]bool)
	set[0] = true

	for range m {
		tmp := make(map[int]bool)
		for k := range set {
			tmp[k+1] = true
			tmp[k+5] = true
			tmp[k+10] = true
			tmp[k+50] = true
		}
		set = tmp
	}

	n -= m
	return len(set) + n*49
}
