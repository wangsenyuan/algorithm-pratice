package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	words := make([]string, n)
	for i := range n {
		words[i] = readString(reader)
	}
	return solve(words)
}

type hp65 struct{ sort.IntSlice }

func (h hp65) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp65) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (hp65) Pop() (_ any)         { return }

func solve(words []string) int {

	type node struct {
		son [26]*node
		end bool
		hp  *hp65
	}

	root := &node{end: true}
	root.hp = &hp65{}

	// hs := map[*node]*hp65{root: {}}
	var ans int
	for _, s := range words {
		ans += len(s)
		o := root
		for _, b := range s {
			b -= 'a'
			if o.son[b] == nil {
				o.son[b] = &node{}
				o.son[b].hp = &hp65{}
			}
			o = o.son[b]
		}
		o.end = true
		o.hp.IntSlice = []int{len(s)}
	}

	var f func(*node, int)
	f = func(o *node, d int) {
		for _, son := range o.son[:] {
			if son != nil {
				f(son, d+1)
				for _, v := range son.hp.IntSlice {
					heap.Push(o.hp, v)
				}
			}
		}
		if !o.end {
			h := o.hp
			ans -= h.IntSlice[0] - d
			h.IntSlice[0] = d
			heap.Fix(h, 0)
		}
	}
	f(root, 0)

	return ans
}
