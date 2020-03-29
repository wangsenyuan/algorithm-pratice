package p1395

import "testing"

func TestSample1(t *testing.T) {
	undergroundSystem := Constructor()

	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	res := undergroundSystem.GetAverageTime("Paradise", "Cambridge") // return 14.0. There was only one travel from "Paradise" (at time 8) to "Cambridge" (at time 22)

	if res != 14.0 {
		t.Fatalf("expect 14.0, but got %f", res)
	}

	res = undergroundSystem.GetAverageTime("Leyton", "Waterloo") // return 11.0. There were two travels from "Leyton" to "Waterloo", a customer with id=45 from time=3 to time=15 and a customer with id=27 from time=10 to time=20. So the average time is ( (15-3) + (20-10) ) / 2 = 11.0

	if res != 11.0 {
		t.Fatalf("expect 11.0, but got %f", res)
	}

	undergroundSystem.CheckIn(10, "Leyton", 24)
	res = undergroundSystem.GetAverageTime("Leyton", "Waterloo") // return 11.0
	if res != 11.0 {
		t.Fatalf("expect 11.0, but got %f", res)
	}
	undergroundSystem.CheckOut(10, "Waterloo", 38)
	res = undergroundSystem.GetAverageTime("Leyton", "Waterloo") // return 12.0
	if res != 12.0 {
		t.Fatalf("expect 12.0, but got %f xxx", res)
	}
}
