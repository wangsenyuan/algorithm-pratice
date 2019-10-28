package p1239

import (
	"testing"
)

func runSample(t *testing.T, arr []string, expect int) {
	res := maxLength(arr)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []string{"un", "iq", "ue"}
	runSample(t, arr, 4)
}
