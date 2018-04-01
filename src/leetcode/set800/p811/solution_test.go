package p811

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, cpdomains []string, expect []string) {
	res := subdomainVisits(cpdomains)
	sort.Strings(expect)
	sort.Strings(res)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample %v, expect %v, but got %v", cpdomains, expect, res)
	}
}

func TestSample1(t *testing.T) {
	cpdomains := []string{"9001 discuss.leetcode.com"}
	expect := []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"}
	runSample(t, cpdomains, expect)
}

func TestSample2(t *testing.T) {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	expect := []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}
	runSample(t, cpdomains, expect)
}
