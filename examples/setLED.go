/*

	Usage:

		./setLED 1
		./setLED 2

	Will flash the LED passed in as the argument, 0 through 17, useful for mapping the LEDs

*/

package main

import (
	"fmt"
	"github.com/wjessop/go-piglow"
	"os"
	"strconv"
	"time"
)

var p *piglow.Piglow

func main() {
	var err error
	p, err = piglow.NewPiglow()
	if err != nil {
		panic(err)
	}

	l, _ := strconv.ParseInt(os.Args[1], 10, 8)
	fmt.Println("Flashing LED ", l)
	flashLED(int8(l))
}

func flashLED(led int8) {
	for i := 0; i < 4; i++ {
		setLED(led, true)
		setLED(led, false)
	}
}

func setLED(led int8, state bool) {
	brightness := 0
	if state == true {
		brightness = 50
	}

	p.SetLED(led, uint8(brightness))
	err := p.Apply()
	if err != nil { // Apply the changes
		panic(err)
	}
	time.Sleep(200 * time.Millisecond)
}
