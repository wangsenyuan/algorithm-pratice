package main

import "fmt"

type StringIterator struct {
	src  string
	pos  int
	next int
	cnt  int
	used int
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func count(str string, pos int) (int, int) {
	if pos >= len(str)-1 {
		return -1, len(str)
	}

	cnt := 0
	i := pos + 1
	for ; i < len(str) && isDigit(str[i]); i++ {
		cnt = cnt*10 + int(str[i]-'0')
	}

	return cnt, i
}

func Constructor(compressedString string) StringIterator {
	cnt, i := count(compressedString, 0)
	return StringIterator{compressedString, 0, i, cnt, 0}
}

func (this *StringIterator) Next() byte {
	if this.pos == len(this.src) {
		return ' '
	}
	this.used++
	rt := this.src[this.pos]

	if this.used == this.cnt {
		this.pos = this.next
		this.used = 0
		cnt, i := count(this.src, this.next)
		this.cnt = cnt
		this.next = i
	}

	return rt
}

func (this *StringIterator) HasNext() bool {
	return this.pos < len(this.src)
}

func main() {
	iter := Constructor("L1e2t1C1o1d1e1")
	for iter.HasNext() {
		fmt.Printf("%s\n", string(iter.Next()))
	}
}
