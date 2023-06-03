package actuator

import (
	"errors"
	"fmt"
	"machine"

	"tinygo.org/x/drivers/servo"
)

func NewActuator(pwm servo.PWM, pin machine.Pin, minSignal int16, maxSignal int16, rangeLimiter uint8) (*Actuator, error) {
	if minSignal >= maxSignal {
		return nil, errors.New("max signal must be greater than min signal")
	}
	if maxSignal > 2000 {
		return nil, errors.New("max signal is too high")
	}
	if minSignal < 1000 {
		return nil, errors.New("min signal is too low")
	}

	s, err := servo.New(pwm, pin)
	if err != nil {
		return nil, err
	}

	return &Actuator{
		minSignal:   minSignal,
		maxSignal:   maxSignal,
		signalRange: int16((float32(maxSignal-minSignal) * float32(rangeLimiter) / 100.0)),
		Servo:       s,
	}, nil
}

// Pass speed from 0 to 100
func (a *Actuator) Speed(value int16) {
	newValue := ((value * a.signalRange) / 100) + a.minSignal
	fmt.Println(newValue)
	a.Servo.SetMicroseconds(newValue)
}
