package p2243

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := digitSum(s, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "11111222223"
	k := 3
	expect := "135"
	runSample(t, s, k, expect)
}
