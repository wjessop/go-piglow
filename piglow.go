package piglow

import (
	"bitbucket.org/gmcbay/i2c"
)

var i2cbus *i2c.I2CBus

const (
	i2c_bus_num byte = 1
	i2c_addr    byte = 0x54 // fixed i2c address of SN3218 ic

	CMD_ENABLE_OUTPUT  byte = 0x00
	CMD_ENABLE_LEDS    byte = 0x13
	CMD_SET_PWM_VALUES byte = 0x01
	CMD_UPDATE         byte = 0x16
)

var white_leds = [3]int8{12, 9, 10}
var blue_leds = [3]int8{14, 4, 11}
var green_leds = [3]int8{3, 5, 13}
var yellow_leds = [3]int8{2, 8, 15}
var orange_leds = [3]int8{1, 7, 16}
var red_leds = [3]int8{0, 6, 17}

type Piglow struct {
	values [18]byte
}

func NewPiglow() (piglow *Piglow, err error) {
	piglow = new(Piglow)
	piglow.values = [18]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	i2cbus, err = i2c.Bus(i2c_bus_num)
	if err != nil {
		return nil, err
	}

	write_err := i2cbus.WriteByte(i2c_addr, CMD_ENABLE_OUTPUT, 0x01)
	if write_err != nil {
		return nil, write_err
	}

	enable_err := i2cbus.WriteByteBlock(i2c_addr, CMD_ENABLE_LEDS, []byte{0xFF, 0xFF, 0xFF})
	if enable_err != nil {
		return nil, enable_err
	}

	return
}

func (p *Piglow) Apply() (err error) {
	err = i2cbus.WriteByteBlock(i2c_addr, CMD_SET_PWM_VALUES, p.values[0:18])
	if err != nil {
		return
	}

	// Tell the i2c device to apply the changes
	err = i2cbus.WriteByte(i2c_addr, CMD_UPDATE, 0xFF)
	if err != nil {
		return
	}

	return err
}

// Set LED n to brightness
// n must be 0 through 17
// brightness must be 0 through 255
func (p *Piglow) SetLED(n int8, brightness uint8) {
	p.values[n] = brightness
}

// Set all LEDs to brightness
func (p *Piglow) SetAll(brightness uint8) {
	for i := 0; i < 18; i++ {
		p.values[i] = brightness
	}
}

/*

	Set colour rings to the given brightness

*/

func (p *Piglow) SetWhite(brightness uint8) {
	for i := 0; i < 3; i++ {
		p.values[white_leds[i]] = brightness
	}
}

func (p *Piglow) SetBlue(brightness uint8) {
	for i := 0; i < 3; i++ {
		p.values[blue_leds[i]] = brightness
	}
}

func (p *Piglow) SetGreen(brightness uint8) {
	for i := 0; i < 3; i++ {
		p.values[green_leds[i]] = brightness
	}
}

func (p *Piglow) SetYellow(brightness uint8) {
	for i := 0; i < 3; i++ {
		p.values[yellow_leds[i]] = brightness
	}
}

func (p *Piglow) SetOrange(brightness uint8) {
	for i := 0; i < 3; i++ {
		p.values[orange_leds[i]] = brightness
	}
}

func (p *Piglow) SetRed(brightness uint8) {
	for i := 0; i < 3; i++ {
		p.values[red_leds[i]] = brightness
	}
}
