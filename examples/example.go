package main

import (
	"github.com/wjessop/go-piglow"
	"log"
	"time"
)

func main() {
	var p *piglow.Piglow
	var err error

	// Create a new Piglow
	p, err = piglow.NewPiglow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	p.SetLED(0, 255) // Set LED 0 to 255 (max brightness)
	p.SetLED(1, 128) // Set LED 1 to half brightness
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}

	time.Sleep(1e9) // Admire your handiwork

	p.SetLED(3, 1) // Set LED 1 to half brightness

	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}

	time.Sleep(1e9) // Admire your handiwork

	p.SetAll(0) // Turn all LEDs off
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}

	time.Sleep(1e7)

	p.SetAll(50) // Turn all LEDs on at 50, max would be 255. Ow, ow my eyes!
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}

	// Fade out
	for i := 50; i >= 0; i-- {
		p.SetAll(uint8(i))
		err = p.Apply()
		if err != nil { // Apply the changes
			log.Fatal("Couldn't apply changes: ", err)
		}
		time.Sleep(5e7)
	}
}
