package main

var temperature int // HL

// ChangeTemp public method has access to private variables in the package
func ChangeTemp(newTemp int) {
	if newTemp < 75 {
		h.thermostat = newTemp
	}
}
