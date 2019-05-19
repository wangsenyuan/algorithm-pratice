package p1047

import "testing"

func runSample(t *testing.T, S string, expect string) {
	res := removeDuplicates(S)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abbaca", "ca")
}

func TestSample2(t *testing.T) {
	runSample(t, "ibfjcaffccadidiaidchakchchcahabhibdcejkdkfbaeeaecdjhajbkfebebfea", "ibfjcdidiaidchakchchcahabhibdcejkdkfbecdjhajbkfebebfea")
}

func TestSample3(t *testing.T) {
	runSample(t, "aaaaaaaaa", "a")
}
