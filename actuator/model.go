package actuator

import "tinygo.org/x/drivers/servo"

type Actuator struct {
	minSignal           int16
	maxSignal           int16
	signalRange         int16
	halfRotationTime    int16
	pulsesPerRevolution int16
	Servo               servo.Servo
}
