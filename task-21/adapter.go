package main

import (
	"fmt"
	"strconv"
)

// Light Class to adapt to the Lighting interface
type Light struct {
	IsOn bool
}

func (l *Light) GetIsOn() bool {
	return l.IsOn
}

func (l *Light) SetIsOn(isOn bool) {
	l.IsOn = isOn
}

//string method to implement the Stringer interface
func (l *Light) String() string {
	return "Light: " + strconv.FormatBool(l.IsOn)
}

// Flashlight interface
type Flashlight interface {
	Light()
}

// LightSwitchAdapter Adapter class for the light switch
type LightSwitchAdapter struct {
	light Light // adaptee object
}

// Light Adapter implementation of the Flashlight interface
func (la *LightSwitchAdapter) Light() {
	light := la.light

	if light.GetIsOn() {
		fmt.Println("Light is on")
	} else {
		light.SetIsOn(true)
		fmt.Println("Light is off")
	}
}

func main() {

	// create a Flashlight interface object
	var lightSwitch Flashlight = &LightSwitchAdapter{
		light: Light{false},
	}

	// use light switch
	lightSwitch.Light()

}
