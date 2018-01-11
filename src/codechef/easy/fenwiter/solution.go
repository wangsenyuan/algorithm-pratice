package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		var a, b, c int
		i := 0
		for bs[i] != ' ' {
			i++
		}
		a = i
		l1 := bs[:a]
		i++
		for bs[i] != ' ' {
			i++
		}
		b = i
		l2 := bs[a+1 : b]
		i++
		for bs[i] != ' ' {
			i++
		}
		c = i
		l3 := bs[b+1 : c]
		i++
		var n int
		readInt(bs, i, &n)
		fmt.Println(solve(l1, l2, l3, n))
	}

}

func solve(l1, l2, l3 []byte, n int) int {
	fn := func(s []byte) (cnt int, suf int) {
		for i := 0; i < len(s); i++ {
			if s[i] == '0' {
				suf = 0
			} else {
				cnt++
				suf++
			}
		}
		return
	}
	cnt1, suf1 := fn(l1)
	cnt2, suf2 := fn(l2)
	cnt3, suf3 := fn(l3)
	var ans = cnt1 + n*cnt2 + cnt3
	if suf3 < len(l3) {
		ans -= suf3 - 1
	} else if suf2 < len(l2) {
		ans -= suf3 + suf2 - 1
	} else {
		ans -= suf3 + n*suf2 + suf1 - 1
	}
	return ans
}
