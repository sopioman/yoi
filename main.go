package main

import (
	"embed"
	"log"
	"machine"

	"github.com/sopioman/yoi/actuator"
	"github.com/sopioman/yoi/feedback"
	"github.com/sopioman/yoi/funscript"
)

var (
	//go:embed data/*
	res embed.FS
)

func main() {
	// starting blink pattern
	feedback.Blink(2, 300)

	fs, err := funscript.NewFuncscipt(res, "data/run.funscript")
	if err != nil {
		log.Fatal(err)
	}

	//feedback.Blink(5, 100)
	fs.Process()

	motor, err := actuator.NewActuator(machine.PWM0, machine.D9, 1050, 1950, fs.Range, 217, 769)
	if err != nil {
		log.Fatal(err)
	}

	motor.Speed(50)

	feedback.Blink(2, 500) // program done
}
