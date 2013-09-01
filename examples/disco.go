package main

import (
	"github.com/wjessop/go-piglow"
	"log"
	"math/rand"
	"time"
)

var max_brightness uint8 = 150

// Cycling from 0 to 255 takes too long, even with no delay so we use intermediate brightness values
var intermediate_brightness_values = [...]int{0x01, 0x02, 0x04, 0x08, 0x10, 0x18, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80, 0x90, 0xA0, 0xC0, 0xE0, 0xFF}
var delay time.Duration = 500 * time.Nanosecond
var p *piglow.Piglow

func main() {
	var err error

	// Create a new Piglow
	p, err = piglow.NewPiglow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	for {
		cycleColoursOut(10)
		allFade(2)
		spinFade(1)
		displayProgressiveValues(2, true)
		cycleColoursOut(2)
		displayProgressiveValues(2, true)
		populateRandomly(75)
		displayProgressiveValues(2, false)
	}
}

func populateRandomly(times int) {
	zero()
	for t := 0; t < times; t++ {
		led_states := [...]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
		for i := 0; i < 18; i++ {
			led := rand.Intn(18)
			if led_states[led] {
				led_states[led] = false
				for b := (len(intermediate_brightness_values) - 1); b >= 0; b-- {
					p.SetLED(int8(led), uint8(b))
					time.Sleep(10 * time.Nanosecond)
				}
			} else {
				led_states[led] = true
				for b := 0; b < len(intermediate_brightness_values); b++ {
					p.SetLED(int8(led), uint8(b))
					time.Sleep(10 * time.Nanosecond)
				}
			}
			apply()
		}

	}
}

func cycleColoursOut(times int) {
	num_brightnesses := len(intermediate_brightness_values)

	for i := 0; i < times; i++ {
		// Fade in white
		for i := 0; i < len(intermediate_brightness_values); i++ {
			p.SetWhite(uint8(intermediate_brightness_values[num_brightnesses-i-1]))
			apply()
			time.Sleep(delay)
		}

		// Fade in blue, out white
		for i := 0; i < len(intermediate_brightness_values); i++ {
			p.SetBlue(uint8(intermediate_brightness_values[i]))
			p.SetWhite(uint8(intermediate_brightness_values[num_brightnesses-i-1]))
			apply()
			time.Sleep(delay)
		}

		// Fade in green, out blue
		for i := 0; i < len(intermediate_brightness_values); i++ {
			p.SetGreen(uint8(intermediate_brightness_values[i]))
			p.SetBlue(uint8(intermediate_brightness_values[num_brightnesses-i-1]))
			apply()
			time.Sleep(delay)
		}

		// Fade in yellow, out green
		for i := 0; i < len(intermediate_brightness_values); i++ {
			p.SetYellow(uint8(intermediate_brightness_values[i]))
			p.SetGreen(uint8(intermediate_brightness_values[num_brightnesses-i-1]))
			apply()
			time.Sleep(delay)
		}

		// Fade in orange, out yellow
		for i := 0; i < len(intermediate_brightness_values); i++ {
			p.SetOrange(uint8(intermediate_brightness_values[i]))
			p.SetYellow(uint8(intermediate_brightness_values[num_brightnesses-i-1]))
			apply()
			time.Sleep(delay)
		}

		// Fade in red, out orange
		for i := 0; i < len(intermediate_brightness_values); i++ {
			p.SetRed(uint8(intermediate_brightness_values[i]))
			p.SetOrange(uint8(intermediate_brightness_values[num_brightnesses-i-1]))
			apply()
			time.Sleep(delay)
		}

		// Fade out Red
		for i := 0; i < len(intermediate_brightness_values); i++ {
			p.SetRed(uint8(intermediate_brightness_values[num_brightnesses-i-1]))
			apply()
			time.Sleep(delay)
		}
	}
}

func allFade(times int) {
	zero()

	for t := 0; t < times; t++ {
		// Fade in
		for i := 0; i <= int(max_brightness); i++ {
			p.SetAll(uint8(i))
			apply()
			time.Sleep(delay)
		}

		// Fade out
		for i := int(max_brightness); i >= 0; i-- {
			p.SetAll(uint8(i))
			apply()
			time.Sleep(delay)
		}
	}
}

func spinFade(times int) {
	for x := 0; x < times; x++ {
		for i := 0; i < len(intermediate_brightness_values); i++ {
			for t := 0; t < 3; t++ {
				p.SetTentacle(t, uint8(intermediate_brightness_values[i]))
				apply()
				time.Sleep(40 * time.Millisecond)
				p.SetAll(0)
			}
		}

		for i := (len(intermediate_brightness_values) - 1); i > 0; i-- {
			for t := 0; t < 3; t++ {
				p.SetTentacle(t, uint8(intermediate_brightness_values[i]))
				apply()
				time.Sleep(40 * time.Millisecond)
				p.SetAll(0)
			}
		}
	}
}

func displayProgressiveValues(times int, direction bool) {
	for t := 0; t < times; t++ {
		max_values_for_tentacle := 6 * len(intermediate_brightness_values)
		for i := 0; i <= max_values_for_tentacle; i++ {
			p.DisplayValueOnTentacle(0, float64(i), float64(max_values_for_tentacle), max_brightness, direction)
			p.DisplayValueOnTentacle(1, float64(i), float64(max_values_for_tentacle), max_brightness, direction)
			p.DisplayValueOnTentacle(2, float64(i), float64(max_values_for_tentacle), max_brightness, direction)
			apply()
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func zero() {
	p.SetAll(0)
}

func apply() {
	err := p.Apply()
	if err != nil {
		log.Fatal("Couldn't apply changes: ", err)
	}
}
