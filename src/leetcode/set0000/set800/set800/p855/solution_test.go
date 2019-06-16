package p855

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, ops []int, expect []int) {
	res := make([]int, 0, len(expect))

	room := Constructor(N)

	for _, op := range ops {
		if op < 0 {
			seat := room.Seat()
			res = append(res, seat)
		} else {
			room.Leave(op)
		}
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v, expect %v, but got %v", N, ops, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 10
	ops := []int{
		-1, -1, -1, -1, 4, -1,
	}
	expect := []int{
		0, 9, 4, 2, 5,
	}
	runSample(t, N, ops, expect)
}
