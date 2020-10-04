package p1611

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 {
		if this.big == 0 {
			return false
		}
		this.big--
	} else if carType == 2 {
		if this.medium == 0 {
			return false
		}
		this.medium--
	} else {
		if this.small == 0 {
			return false
		}
		this.small--
	}
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
