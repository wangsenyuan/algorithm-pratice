package main

import (
	"sort"
	"bufio"
	"os"
	"fmt"
)

var num []int

func init() {
	num = make([]int, 10000000)
	masks := make([]int, 10000000)
	b2i := make([]int, 3000)

	for i := 0; i < 10; i++ {
		b2i[1<<uint(i)] = i
	}

	sz := 0
	for i := 1; i <= 9; i++ {
		num[sz] = i
		masks[sz] = ((1 << 10) - 1) ^ (1 << uint(i))
		sz++
	}

	for k := 0; k < sz; k++ {
		rest := masks[k]
		for rest > 0 {
			bt := rest & -rest
			i := b2i[bt]
			num[sz] = num[k]*10 + i
			masks[sz] = masks[k] ^ (1 << uint(i))
			sz++
			rest ^= bt
		}
	}
	num = num[:sz]
}

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readInt64(bs []byte, from int, val *int64) int {
	var tmp int64
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int64(bs[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)
	for i := 0; i < t; i++ {
		var L, R int64
		scanner.Scan()
		bs := scanner.Bytes()
		x := readInt64(bs, 0, &L)
		readInt64(bs, x+1, &R)
		res := solve(L, R)
		fmt.Println(res)
	}
}

func solve(L, R int64) int {
	i := sort.Search(len(num), func(i int) bool {
		return int64(num[i]) >= L
	})
	j := sort.Search(len(num), func(j int) bool {
		return int64(num[j]) > R
	})
	return j - i
}
