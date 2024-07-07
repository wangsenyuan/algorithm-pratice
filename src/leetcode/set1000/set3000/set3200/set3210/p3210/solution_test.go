package p3210

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := getEncryptedString(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "dart"
	k := 3
	expect := "tdar"
	runSample(t, s, k, expect)
}
