/*

	A CPU meter. Reads in the Raspberry Pi's load average and
	displays the 1, 5 and 15 minute load average on different
	"tentacles" of the Piglow.

*/

package main

import (
	"fmt"
	"github.com/wjessop/go-piglow"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

const (
	max_load       = 1.0 // The load at which we display a full bar of LEDs. Tweak for your system
	max_brightness = 10  // Any more is eye-searingly painful to look at
)

var p *piglow.Piglow
var load1_leds = [6]int8{12, 14, 3, 2, 1, 0}
var load2_leds = [6]int8{9, 4, 5, 8, 7, 6}
var load3_leds = [6]int8{10, 11, 13, 15, 16, 17}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			p.SetAll(0)
			err := p.Apply()
			if err != nil {
				panic(err)
			}
			os.Exit(0)
		}
	}()

	var err error

	p, err = piglow.NewPiglow()
	if err != nil {
		panic(err)
	}

	for {
		load_data, err := ioutil.ReadFile("/proc/loadavg")
		if err != nil {
			panic(err)
		}

		vals := strings.Split(string(load_data), " ")

		fmt.Printf("Load avg: %v %v %v\n", vals[0], vals[1], vals[2])
		p.DisplayValueOnTentacle(0, intFromLoadStr(vals[0]), max_load, max_brightness, true)
		p.DisplayValueOnTentacle(1, intFromLoadStr(vals[1]), max_load, max_brightness, true)
		p.DisplayValueOnTentacle(2, intFromLoadStr(vals[2]), max_load, max_brightness, true)

		err = p.Apply()
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func zeroLEDs(leds [6]int8) {
	for i := 0; i < 6; i++ {
		p.SetLED(leds[i], 0)
	}
}

func intFromLoadStr(str string) (i float64) {
	i, err := strconv.ParseFloat(str, 32)
	if err != nil {
		panic(err)
	}

	return
}
