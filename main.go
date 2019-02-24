// Package main implements a program that will run a hardware PWM channel through
// a sequence of duty cycles to demonstrate programmatic PWM control
//
// Pass '-help' to get commandline help
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/chmorgan/go-pwm-example/pwm"
)

// CommandLineOptions contains all of the potential command line flags that can be specified
type CommandLineOptions struct {
	channel     *int
	frequencyHz *float64
}

var commandLineOptions CommandLineOptions

func parseCommandLineOptions(c *CommandLineOptions) {
	c.channel = flag.Int("channel", 0, "PWM channel (0, 1, 2, or 3)")
	c.frequencyHz = flag.Float64("frequency", 100.0, "Frequency in Hz")
	flag.Parse()
}

func main() {
	unsugardedLogger, _ := zap.NewDevelopment()
	logger := unsugardedLogger.Sugar()

	parseCommandLineOptions(&commandLineOptions)


	fmt.Printf("pwm channel: %d, frequencyHz: %f\n", *commandLineOptions.channel, *commandLineOptions.frequencyHz)
	fmt.Printf("Pass '-help' as an argument to see usage\n\n")

	pwm, err := pwm.New(logger, *commandLineOptions.channel, *commandLineOptions.frequencyHz)
	if err != nil {
		fmt.Println("pwm.New failed ", err)
		os.Exit(1)
	}

	// set a 50% duty cycle
	dutyCycle := 50.0
	fmt.Printf("Setting duty cycle to %.1f%%\n", dutyCycle)
	pwm.SetDutyCycle(dutyCycle)

	time.Sleep(time.Second * 3)

	// set a 100% duty cycle
	dutyCycle = 100.0
	fmt.Printf("Setting duty cycle to %.1f%%\n", dutyCycle)
	pwm.SetDutyCycle(dutyCycle)

	time.Sleep(time.Second * 3)

	// disable pwm output
	dutyCycle = 0.0
	fmt.Printf("Setting duty cycle to %.1f%%\n", dutyCycle)
	pwm.SetDutyCycle(dutyCycle)
}
