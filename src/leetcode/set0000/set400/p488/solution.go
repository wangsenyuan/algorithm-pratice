package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(findMinStep("WRRBBW", "RB"))
	fmt.Println(findMinStep("WWRRBBWW", "WRBRW"))
	fmt.Println(findMinStep("G", "GGGGG"))
	fmt.Println(findMinStep("RBYYBBRRB", "YRBGB"))
}

func findMinStep(board string, hand string) int {
	h := []byte(hand)
	sort.Sort(ByteSlice(h))
	b := []byte(board)

	var help func(b, h []byte) int

	help = func(b, h []byte) int {
		if len(b) == 0 {
			return 0
		}

		if len(h) == 0 {
			return 6
		}

		res := 6
		for i := 0; i < len(h); i++ {
			var j int
			for j < len(b) {
				k, found := find(b, j, h[i])
				if !found {
					break
				}

				if k < len(b)-1 && b[k] == b[k + 1] {
					B := shrink(concat(b[:k], b[k + 2:]))
					if len(B) == 0 {
						return 1
					}

					H := concat(h[:i], h[i + 1:])

					res = min(res, help(B, H)+1)

					k++
				} else if i > 0 && h[i] == h[i - 1] {
					B := shrink(concat(b[:k], b[k + 1:]))
					if len(B) == 0 {
						return 2
					}
					H := concat(h[:i - 1], h[i + 1:])
					res = min(res, help(B, H)+2)
				}

				j = k + 1
			}
		}
		return res
	}

	res := help(b, h)

	if res == 6 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func shrink(src []byte) []byte {
	dst := make([]byte, 0, len(src))
	var j int
	for i := 1; i <= len(src); i++ {
		if i < len(src) && src[i] == src[i - 1] {
			continue
		}

		if i-j < 3 {
			for k := j; k < i; k++ {
				dst = append(dst, src[k])
			}
		}
		j = i
	}

	if len(src) == len(dst) {
		return dst
	}

	return shrink(dst)
}

func find(src []byte, from int, x byte) (int, bool) {
	for i := from; i < len(src); i++ {
		if src[i] == x {
			return i, true
		}
	}
	return -1, false
}

func concat(a, b []byte) []byte {
	c := make([]byte, len(a))
	copy(c, a)
	return append(c, b...)
}

type ByteSlice []byte

func (bs ByteSlice) Len() int {
	return len(bs)
}

func (bs ByteSlice) Less(i, j int) bool {
	return bs[i] < bs[j]
}

func (bs ByteSlice) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
