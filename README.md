# piglow

go-piglow is a small lib for controlling the [Piglow](http://shop.pimoroni.com/products/piglow) from Go on your Raspberry Pi.

See the [Piglow github page](https://github.com/pimoroni/piglow) for details on how to set up your Raspberry Pi for this to work.

## Example

````go
package main

import (
	"github.com/wjessop/go-piglow"
	"log"
)

func main() {
	var p *piglow.Piglow
	var err error

	// Create a new Piglow
	p, err = piglow.NewPiglow(); if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	p.SetLED(0, 255) // Set LED 0 to 255 (max brightness)
	p.SetLED(1, 128) // Set LED 1 to half brightness
	err = p.Apply(); if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
}
````

Cross compile for the Raspberry pi with:

````GOOS=linux GOARM=6 GOARCH=arm go build````

See examples/example.go for a list of functions you can use. The other files in examples/* are more complex examples of how to use the lib.

## Notes

- The LEDs aren't in order, experiment to find out which is which
- Setting all LEDs to 255 will probably hurt your eyes, I don't recommend it
- The lib isn't thread-safe (It could be, but I don't see the point) , so only create one instance of Piglow

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Author

* Will Jessop, @will_j, will@willj.net
