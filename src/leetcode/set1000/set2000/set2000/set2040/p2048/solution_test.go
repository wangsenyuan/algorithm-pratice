package p2048

import (
	"testing"
)

func runSample(t *testing.T, n int, expect int) {
	res := nextBeautifulNumber(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 22)
}

func TestSample2(t *testing.T) {
	runSample(t, 1000, 1333)
}

func TestSample3(t *testing.T) {
	runSample(t, 3000, 3133)
}

func TestSample4(t *testing.T) {
	runSample(t, 188, 212)
}

func TestSample5(t *testing.T) {
	runSample(t, 16407, 22333)
}
