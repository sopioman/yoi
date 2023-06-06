package feedback

import (
	"machine"
	"time"
)

var (
	led machine.Pin
)

func init() {
	led = machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func Blink(count int, ms int32) {
	for i := 0; i < count; i++ {
		led.High()
		time.Sleep(time.Millisecond * time.Duration(ms))
		led.Low()
		time.Sleep(time.Millisecond * time.Duration(ms))
	}

}
