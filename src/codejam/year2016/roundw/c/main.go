package main

import "fmt"

func main() {
	var t = 0
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var c, v, l = 0, 0, 0
		fmt.Scanf("%d %d %d", &c, &v, &l)
		r := play(c, v, l)
		fmt.Printf("Case #%d: %d\n", i, r)
	}
}

const (
	mask int64 = 1000000007
	V    byte  = 'V'
	C    byte  = 'C'
)

func updateSlice(slice [][]byte, a []byte) [][]byte {
	if len(slice)+1 > cap(slice) {
		newSlice := make([][]byte, len(slice), 2*len(slice))
		copy(newSlice, slice)
		slice = newSlice
	}
	return append(slice, a)
}

func asSlice(a byte) []byte {
	return []byte{a}
}

func template(l int, sliceLength int) [][]byte {
	result := make([][]byte, 0, 1)
	if l == 0 {
		emptySlice := make([]byte, 0, sliceLength)
		return updateSlice(result, emptySlice)
	}

	subTemplate := template(l-1, sliceLength)

	for _, sub := range subTemplate {
		if len(sub) == 0 {
			result = updateSlice(result, asSlice(V))
			continue
		}
		result = updateSlice(result, append(asSlice(V), sub...))
		h := sub[0]
		if h == V {
			result = updateSlice(result, append(asSlice(C), sub...))
		}
	}
	return result
}

func calculate(t []byte, c, v int) int64 {
	var p int64 = 1
	for _, x := range t {
		if x == V {
			p *= int64(v)
		} else {
			p *= int64(c)
		}
		p %= mask
	}

	return p
}

func play(c, v, l int) int64 {
	ts := template(l, l)
	var sum int64 = 0

	for _, t := range ts {
		x := calculate(t, c, v) % mask
		sum += x
		sum %= mask
	}

	return sum % mask
}
