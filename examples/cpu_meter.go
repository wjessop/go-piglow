/*

	A CPU meter. Reads in the Raspberry Pi's load average and
	displays the 1, 5 and 15 minute load average on different
	"arms" of the Piglow.

*/

package main

import (
	"fmt"
	"github.com/wjessop/go-piglow"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	max_load       = 1  // The load at which we display a full bar of LEDs. Tweak for your system
	max_brightness = 10 // Any more is eye-searingly painful to look at
)

var p *piglow.Piglow
var load1_leds = [6]int8{12, 14, 3, 2, 1, 0}
var load2_leds = [6]int8{9, 4, 5, 8, 7, 6}
var load3_leds = [6]int8{10, 11, 13, 15, 16, 17}

func main() {
	var err error

	p, err = piglow.NewPiglow()
	if err != nil {
		panic(err)
	}

	for {
		// 0.00 0.01 0.05 2/64 1269
		load_data, err := ioutil.ReadFile("/proc/loadavg")
		if err != nil {
			panic(err)
		}

		vals := strings.Split(string(load_data), " ")

		fmt.Printf("Load avg: %v %v %v\n", vals[0], vals[1], vals[2])
		setLEDvalues(intFromLoadStr(vals[0]), load1_leds)
		setLEDvalues(intFromLoadStr(vals[1]), load2_leds)
		setLEDvalues(intFromLoadStr(vals[2]), load3_leds)

		err = p.Apply()
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func setLEDvalues(v uint16, leds [6]int8) {
	zeroLEDs(leds)

	prop := float64(v) / max_brightness
	full, partial := math.Modf(prop)
	for i := 0; i < int(full); i++ {
		p.SetLED(leds[i], max_brightness)
	}

	if partial == 0 {
		return
	}

	partial_brightness := math.Floor((max_brightness * partial) + 0.5)
	p.SetLED(leds[int(full)], uint8(partial_brightness))
}

func zeroLEDs(leds [6]int8) {
	for i := 0; i < 6; i++ {
		p.SetLED(leds[i], 0)
	}
}

func intFromLoadStr(str string) (i uint16) {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		panic(err)
	}

	if f >= max_load {
		i = 6 * max_brightness
	} else {
		l := (6 * max_brightness) * f
		i = uint16(math.Floor(l + 0.5))
	}
	return
}
