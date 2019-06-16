package p744

import "testing"

func TestSamples(t *testing.T) {
	samples := []struct {
		letters []byte
		target  byte
		expect  byte
	}{
		{[]byte{'c', 'f', 'g'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'c', 'f', 'j'}, 'g', 'j'},
		{[]byte{'c', 'f', 'j'}, 'j', 'c'},
		{[]byte{'c', 'f', 'j'}, 'k', 'c'},
	}

	for _, sample := range samples {
		res := nextGreatestLetter(sample.letters, sample.target)
		if res != sample.expect {
			t.Errorf("sample %v should give answer %v, but got %v", sample.letters, sample.expect, res)
		}
	}
}
