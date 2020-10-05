package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s, t string
	//fmt.Scanf("%s", &s)
	//fmt.Scanf("%s", &t)
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	t, _ = reader.ReadString('\n')
	res := solve(s, t)
	fmt.Println(len(res))
	for _, x := range res {
		fmt.Println(x)
	}
}

func solve(s, t string) []string {
	dx := int(t[0]) - int(s[0])
	dy := int(t[1]) - int(s[1])
	res := make([]string, 0, 8)
	if dx != 0 && dy != 0 {
		cnt := min(abs(dx), abs(dy))
		var a = "R"
		if dx < 0 {
			a = "L"
			dx += cnt
		} else {
			dx -= cnt
		}
		var b = "U"
		if dy < 0 {
			b = "D"
			dy += cnt
		} else {
			dy -= cnt
		}
		for cnt > 0 {
			res = append(res, a+b)
			cnt--
		}
	}

	if dx != 0 {
		a := "R"
		if dx < 0 {
			a = "L"
		}
		cnt := abs(dx)
		for cnt > 0 {
			res = append(res, a)
			cnt--
		}
	}
	if dy != 0 {
		b := "U"
		if dy < 0 {
			b = "D"
		}
		cnt := abs(dy)
		for cnt > 0 {
			res = append(res, b)
			cnt--
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
