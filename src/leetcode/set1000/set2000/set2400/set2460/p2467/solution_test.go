package p2467

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, message string, limit int, expect []string) {
	res := splitMessage(message, limit)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	message := "this is really a very awesome message"
	limit := 9
	expect := []string{"thi<1/14>", "s i<2/14>", "s r<3/14>", "eal<4/14>", "ly <5/14>", "a v<6/14>", "ery<7/14>", " aw<8/14>", "eso<9/14>", "me<10/14>", " m<11/14>", "es<12/14>", "sa<13/14>", "ge<14/14>"}
	runSample(t, message, limit, expect)
}

func TestSample2(t *testing.T) {
	message := "short message"
	limit := 15
	expect := []string{"short mess<1/2>", "age<2/2>"}
	runSample(t, message, limit, expect)
}
