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

	// Fade in white
	for i := 0; i <= 50; i++ {
		p.SetWhite(uint8(i))
		err = p.Apply()
		if err != nil { // Apply the changes
			log.Fatal("Couldn't apply changes: ", err)
		}
		time.Sleep(5e7)
	}

	// Fade in blue, out white
	for i := 0; i <= 50; i++ {
		p.SetBlue(uint8(i))
		p.SetWhite(uint8(50 - i))
		err = p.Apply()
		if err != nil { // Apply the changes
			log.Fatal("Couldn't apply changes: ", err)
		}
		time.Sleep(5e7)
	}

	// Fade in green, out blue
	for i := 0; i <= 50; i++ {
		p.SetGreen(uint8(i))
		p.SetBlue(uint8(50 - i))
		err = p.Apply()
		if err != nil { // Apply the changes
			log.Fatal("Couldn't apply changes: ", err)
		}
		time.Sleep(5e7)
	}

	// Fade in yellow, out green
	for i := 0; i <= 50; i++ {
		p.SetYellow(uint8(i))
		p.SetGreen(uint8(50 - i))
		err = p.Apply()
		if err != nil { // Apply the changes
			log.Fatal("Couldn't apply changes: ", err)
		}
		time.Sleep(5e7)
	}

	// Fade in orange, out yellow
	for i := 0; i <= 50; i++ {
		p.SetOrange(uint8(i))
		p.SetYellow(uint8(50 - i))
		err = p.Apply()
		if err != nil { // Apply the changes
			log.Fatal("Couldn't apply changes: ", err)
		}
		time.Sleep(5e7)
	}

	// Fade in red, out orange
	for i := 0; i <= 50; i++ {
		p.SetRed(uint8(i))
		p.SetOrange(uint8(50 - i))
		err = p.Apply()
		if err != nil { // Apply the changes
			log.Fatal("Couldn't apply changes: ", err)
		}
		time.Sleep(5e7)
	}

	p.SetAll(0) // Turn them all off

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
