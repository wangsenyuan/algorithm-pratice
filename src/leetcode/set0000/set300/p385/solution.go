package main

import (
	"fmt"
	"strconv"
	"strings"
)

type NestedInteger struct {
	val      int
	children []NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return len(n.children) == 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.val
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.children = append(n.children, elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []NestedInteger {
	return n.children
}

func next(s string, i int, level int) int {
	if s[i] == ']' && level == 1 {
		return i
	}
	if s[i] == ']' {
		return next(s, i+1, level-1)
	}
	if s[i] == '[' {
		return next(s, i+1, level+1)
	}
	return next(s, i+1, level)
}

func split(s string) []string {
	i, j := 0, 0
	var rs []string
	for j < len(s) {
		if s[j] == ',' {
			rs = append(rs, s[i:j])
			i = j + 1
			j++
			continue
		}
		if s[j] == '[' {
			k := next(s, j, 0)
			rs = append(rs, s[j:k+1])
			i = k + 2
			j = k + 2
			continue
		}
		j++
	}
	if i < len(s) {
		rs = append(rs, s[i:len(s)])
	}
	return rs
}

func deserialize(s string) *NestedInteger {
	x := &NestedInteger{}
	if !strings.HasPrefix(s, "[") {
		val, _ := strconv.Atoi(s)
		x.SetInteger(val)
		return x
	}

	for _, part := range split(s[1 : len(s)-1]) {
		y := deserialize(part)
		x.Add(*y)
	}
	return x
}

func main() {
	//x := deserialize("123")
	//fmt.Printf("x is integer: %v\n", x.IsInteger())
	y := deserialize("[123,[456,[789]]]")
	output(y)
}

func output(n *NestedInteger) {
	if n == nil {
		return
	}

	if n.IsInteger() {
		fmt.Printf("%d", n.GetInteger())
		return
	}

	fmt.Print("[")
	for _, x := range n.GetList() {
		output(&x)
	}
	fmt.Print("]")
}
