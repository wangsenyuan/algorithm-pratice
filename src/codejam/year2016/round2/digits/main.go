package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type num struct {
	id  rune
	val int
	wg  map[rune]int
}

var nums = []num{
	num{'Z', 0, groupWord("ZERO")},
	num{'X', 6, groupWord("SIX")},
	num{'G', 8, groupWord("EIGHT")},
	num{'W', 2, groupWord("TWO")},
	num{'U', 4, groupWord("FOUR")},
	num{'F', 5, groupWord("FIVE")},
	num{'V', 7, groupWord("SEVEN")},
	num{'T', 3, groupWord("THREE")},
	num{'I', 9, groupWord("NINE")},
	num{'O', 1, groupWord("ONE")}}

func groupWord(word string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range word {
		m[c]++
	}
	return m
}

func play(m map[rune]int, words []num, res []int) []int {
	if len(words) == 0 {
		return res
	}

	n := words[0]
	if y, ok := m[n.id]; ok {
		return play(removeNum(m, n, y), words[1:], appendNum(res, n, y))
	}

	return play(m, words[1:], res)
}

func removeNum(m map[rune]int, n num, cnt int) map[rune]int {
	nm := make(map[rune]int)
	wg := n.wg
	for c, x := range m {
		if y, ok := wg[c]; ok {
			if x > cnt*y {
				nm[c] = x - cnt*y
			}
		} else {
			nm[c] = x
		}
	}
	return nm
}

func appendNum(res []int, n num, cnt int) []int {
	for len(res)+cnt > cap(res) {
		slice := make([]int, len(res), 2*len(res))
		copy(slice, res)
		res = slice
	}

	for i := 0; i < cnt; i++ {
		res = append(res, n.val)
	}

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.Trim(line, "\n "))
	//fmt.Println(line)
	//fmt.Println(t)

	for i := 1; i <= t; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.Trim(line, "\n ")
		xs := play(groupWord(line), nums, make([]int, 0, len(line)))
		sort.Ints(xs)
		fmt.Printf("Case #%d: %s\n", i, concat(xs, ""))
	}
}

func concat(ints []int, sep string) string {
	var buffer bytes.Buffer
	_sep := ""
	for _, x := range ints {
		buffer.WriteString(_sep)
		buffer.WriteString(strconv.Itoa(x))
		_sep = sep
	}
	return buffer.String()
}
