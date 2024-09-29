package p3307

import "testing"

func runSample(t *testing.T, k int, ops []int, expect byte) {
	res := kthCharacter(int64(k), ops)

	if res != expect {
		t.Fatalf("Sample expect %c, but got %c", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 5
	ops := []int{0, 0, 0}
	expect := 'a'
	runSample(t, k, ops, byte(expect))
}

func TestSample2(t *testing.T) {
	k := 10
	ops := []int{0, 1, 0, 1}
	expect := 'b'
	runSample(t, k, ops, byte(expect))
}
