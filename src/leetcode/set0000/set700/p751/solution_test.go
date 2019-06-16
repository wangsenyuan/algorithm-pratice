package p751

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, ip string, r int, expect []string) {
	res := ipToCIDR(ip, r)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %s, %d should give %v, but got %v", ip, r, expect, res)
	}
}

func TestSample1(t *testing.T) {
	ip := "255.0.0.7"
	n := 10
	expect := []string{"255.0.0.7/32", "255.0.0.8/29", "255.0.0.16/32"}
	runSample(t, ip, n, expect)
}
