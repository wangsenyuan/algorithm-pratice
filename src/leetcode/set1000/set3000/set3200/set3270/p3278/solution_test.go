package p3278

import "testing"

func runSample(t *testing.T, date, expect string) {
	res := convertDateToBinary(date)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "2080-02-29", "100000100000-10-11101")
}
