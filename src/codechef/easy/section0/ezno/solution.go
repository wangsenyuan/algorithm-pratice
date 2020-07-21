package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var ans []map[int]int

func mod(divident string, divisor int, base int) int {
	var result int

	for i := 0; i < len(divident); i++ {
		result = (result*base + toInt(divident[i])) % divisor
	}

	return result
}

func init() {
	ans = make([]map[int]int, 17)

	for i := 0; i < 17; i++ {
		ans[i] = make(map[int]int)
	}

	var construct func(base int, current string, length int)

	construct = func(base int, current string, length int) {
		ans[base][length]++
		// fmt.Printf("[debug] %d: %s\n", base, current)
		for x := 0; x < base; x++ {
			next := current + toStr(x)
			if mod(next, length+1, base) == 0 {
				construct(base, next, length+1)
			}
		}
	}

	for base := 2; base <= 16; base++ {
		for x := 1; x < base; x++ {
			construct(base, toStr(x), 1)
		}
	}
	// fmt.Println("base length count")
	// fmt.Printf("%v\n", ans)
	for base := 2; base <= 16; base++ {
		mp := ans[base]
		ks := make([]int, 0, len(mp))
		for k := range mp {
			ks = append(ks, k)
		}
		sort.Ints(ks)
		max := ks[len(ks)-1]
		for i := 1; i <= max; i++ {
			fmt.Printf("%d %d %d", base, i, mp[i])
		}
	}
}

func toStr(x int) string {
	if x < 10 {
		return strconv.Itoa(x)
	}
	return string('A' + (x - 10))
}

func toInt(x byte) int {
	if x <= '9' && x >= '0' {
		return int(x - '0')
	}
	return 10 + int(x-'A')
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		var n, d int
		x := readInt(scanner.Bytes(), 0, &n)
		readInt(scanner.Bytes(), x+1, &d)
		fmt.Println(solve(n, d))
	}
}

func solve(n, d int) int {
	if d >= len(ans[n]) {
		return 0
	}
	return ans[n][d]
}
