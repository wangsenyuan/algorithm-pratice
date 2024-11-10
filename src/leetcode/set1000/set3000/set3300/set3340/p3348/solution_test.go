package p3348

import (
	"testing"
)

func runSample(t *testing.T, num string, x int, expect string) {
	res := smallestNumber(num, int64(x))

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := "1234"
	x := 256
	expect := "1488"
	runSample(t, num, x, expect)
}

func TestSample2(t *testing.T) {
	num := "12355"
	x := 50
	expect := "12355"
	runSample(t, num, x, expect)
}

func TestSample3(t *testing.T) {
	num := "11111"
	x := 26
	expect := "-1"
	runSample(t, num, x, expect)
}

func TestSample4(t *testing.T) {
	num := "4093"
	x := 180
	expect := "4159"
	runSample(t, num, x, expect)
}
