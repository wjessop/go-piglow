package main

import (
	"github.com/wjessop/go-piglow"
)

func main() {
	var p *piglow.Piglow
	var err error

	// Create a new Piglow
	p, err = piglow.NewPiglow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	// Set LED to brightness 10
	p.SetLED(0, 10)

	// Set LED to max brightness
	p.SetLED(2, 255)

	// Set all LEDs to brightness 10
	p.SetAll(10)

	// Set the white LEDs to 15
	p.SetWhite(15)

	// Set the red LEDs to 20
	p.SetRed(20)

	// Other functions are available for the other colours.

	// Set all LEDs on tentacle 0 to brightness 15
	p.SetTentacle(0, 15)

	// Set all LEDs on tentacle 2 to brightness 150
	p.SetTentacle(2, 150)

	// Display a value on a tentacle at brightness 10
	// See code comments for more info on parameters
	p.DisplayValueOnTentacle(0, 727.0, 1000.0, uint8(10), true)
}
