package p1624

import "testing"

func TestSample1(t *testing.T) {
	fancy := Constructor()

	fancy.Append(2)
	fancy.AddAll(3)
	fancy.Append(7)
	fancy.MultAll(2)

	res := fancy.GetIndex(0)

	expect(t, res == 10, "should be 10")

	fancy.AddAll(3)
	fancy.Append(10)
	fancy.MultAll(2)

	res = fancy.GetIndex(0)

	expect(t, res == 26, "should be 26")

	res = fancy.GetIndex(1)

	expect(t, res == 34, "should be 34")

	res = fancy.GetIndex(2)

	expect(t, res == 20, "should be 20")

}

func expect(t *testing.T, expr bool, message string) {
	if !expr {
		t.Fatal(message)
	}
}
