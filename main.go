package main

import (
	"log"
	"machine"

	"github.com/sopioman/yoi/actuator"
	"github.com/sopioman/yoi/funscript"
)

func main() {
	fs, err := funscript.NewFuncscipt("run.funscript")
	if err != nil {
		log.Fatal(err)
	}

	//fs.Process()

	motor, err := actuator.NewActuator(machine.PWM0, machine.D9, 1050, 1950, fs.Range)
	if err != nil {
		log.Fatal(err)
	}

	motor.Speed(50)
}
