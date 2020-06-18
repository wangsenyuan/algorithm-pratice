package recoverfrompreoder

import "testing"

func TestSample1(t *testing.T) {
	s := "1-2--3--4-5--6--7"
	node := recoverFromPreorder(s)
	if node == nil {
		t.Errorf("Sample fail")
	}
}
