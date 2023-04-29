package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	chosen := option1
	if option1 > option2 {
		chosen = option2
	}
	return chosen + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	var percentage int
	if age < 3.0 {
		percentage = 80
	} else if age < 10.0 {
		percentage = 70
	} else {
		percentage = 50
	}

	return (originalPrice * float64(percentage)) / 100
}
