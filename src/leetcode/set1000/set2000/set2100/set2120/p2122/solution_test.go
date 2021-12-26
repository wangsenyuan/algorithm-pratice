package p2122

import "testing"

func runSample(t *testing.T, left, right int, expect string) {
	res := abbreviateProduct(left, right)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	left := 1
	right := 4
	expect := "24e0"
	runSample(t, left, right, expect)
}

func TestSample2(t *testing.T) {
	left := 2
	right := 11
	expect := "399168e2"
	runSample(t, left, right, expect)
}


func TestSample3(t *testing.T) {
	left := 999998
	right := 1000000
	expect := "99999...00002e6"
	runSample(t, left, right, expect)
}

func TestSample4(t *testing.T) {
	left := 256
	right := 65535
	expect := "23510...78528e16317"
	runSample(t, left, right, expect)
}
