package p3273

import "testing"

func runSample(t *testing.T, power int, damage []int, health []int, expect int) {
	res := minDamage(power, damage, health)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{4, 5, 6, 8}
	p := 4
	expect := 39
	runSample(t, p, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{80, 79}
	b := []int{86, 13}
	p := 62
	expect := 319
	runSample(t, p, a, b, expect)
}
