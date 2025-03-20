package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if len(res) == 0 {
		fmt.Println("No solution")
	} else {
		fmt.Println(res[0], res[1])
	}
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	s := readString(reader)
	return solve(s)
}

func solve(s string) []int {
	n := len(s)
	cnt := make([]int, n+1)

	var best int

	// v <= 2 * c
	// v + c = m
	// m = v + c <= 3 * c
	// m <= 3 * c
	// m = r - l
	// r - l <= 3 * (cnt[r] - cnt[l])
	// 这里是为了找出最小的l
	// r - 3 * cnt[r] <= l - 3 * cnt[l]
	// 3 * cnt[r] - r >= 3 * cnt[l] - l
	// 所以要找到前面最小的 3 * cnt[l] - l
	// cnt[i] 到i为止有多非元音字符的数量
	tr := NewTree(n + 1)
	// 3 * 0 - 0
	tr.Update(0, 0)

	for i := range n {
		cnt[i+1] = cnt[i]
		if !isVowel(s[i]) {
			cnt[i+1]++
		}
		cur := 3*cnt[i+1] - (i + 1)
		if tr.arr[0] <= cur {
			l := tr.Query(cur)
			best = max(best, i+1-l)
		}
		tr.Update(i+1, cur)
	}

	if best == 0 {
		// 不存在符合条件的子串 v <= 2 * c (比如全部是元音)
		return nil
	}

	ans := []int{best, 0}

	for i := 0; i+best <= n; i++ {
		if 3*(cnt[i+best]-cnt[i]) >= best {
			ans[1]++
		}
	}

	return ans
}

func isVowel(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

type Tree struct {
	arr []int
	n   int
}

const inf = 1 << 60

func NewTree(n int) *Tree {
	arr := make([]int, n*4)
	for i := range 4 * n {
		arr[i] = inf
	}
	return &Tree{arr, n}
}

func (t *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			t.arr[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		t.arr[i] = min(t.arr[i*2+1], t.arr[i*2+2])
	}
	loop(0, 0, t.n-1)
}

func (t *Tree) Query(val int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if t.arr[2*i+1] <= val {
			return loop(2*i+1, l, mid)
		}
		return loop(2*i+2, mid+1, r)
	}
	return loop(0, 0, t.n-1)
}
