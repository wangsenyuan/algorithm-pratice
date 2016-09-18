package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join(readBinaryWatch(2), "\n"))

}

var hours = make(map[int][]string)
var minutes = make(map[int][]string)

func init() {
	for i := 0; i < 12; i++ {
		hours[i] = hour(i)
	}

	for i := 0; i < 60; i++ {
		minutes[i] = minute(i)
	}
}

func readBinaryWatch(num int) []string {
	if num < 0 || num > 10 {
		return nil
	}

	var res []string
	for i := 0; i <= num && i <= 4; i++ {
		for _, h := range hours[i] {
			for _, m := range minutes[num-i] {
				res = append(res, h+":"+m)
			}
		}
	}

	return res
}

func minute(num int) []string {
	return format(num, 60, func(m int) string {
		return fmt.Sprintf("%02d", m)
	})
}

func hour(num int) []string {
	return format(num, 12, func(h int) string {
		return fmt.Sprintf("%d", h)
	})
}

func format(num int, end int, doFmt func(int) string) []string {
	if num == 0 {
		return []string{doFmt(0)}
	}

	var res []string
	for i := 0; i < end; i++ {
		ones := countOne(i)
		if ones == num {
			res = append(res, doFmt(i))
		}
	}

	return res
}

func countOne(num int) int {
	cnt := 0
	for num > 0 {
		if (num & 1) == 1 {
			cnt++
		}
		num = num >> 1
	}

	return cnt
}
