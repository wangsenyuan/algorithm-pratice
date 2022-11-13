package p2469

func convertTemperature(celsius float64) []float64 {
	a := celsius + 273.15
	b := celsius*1.8 + 32
	return []float64{a, b}
}
