package p2266

import "testing"

func runSample(t *testing.T, keys string, expect int) {
	res := countTexts(keys)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	keys := "22233"
	expect := 8
	runSample(t, keys, expect)
}

func TestSample2(t *testing.T) {
	keys := "222222222222222222222222222222222222"
	expect := 82876089
	runSample(t, keys, expect)
}
