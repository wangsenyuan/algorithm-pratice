package p929

import "testing"

func runSample(t *testing.T, emails []string, expect int) {
	res := numUniqueEmails(emails)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", emails, expect, res)
	}
}

func TestSample1(t *testing.T) {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	expect := 2
	runSample(t, emails, expect)
}
