//Package pwm package manages a sysfs exposed pwm pin, providing helpers to configure the pwm signal
// using more commonly understood values such as frequency and duty cycles
//
// NOTE:
//   When changing period, period must be >= duty_ycle or the device will reject the change
package pwm

import (
	"go.uber.org/zap"
	"gobot.io/x/gobot/sysfs"
)

// Pwm provides an interface to a Pwm pin using percentages rather than nanosecond values
type Pwm interface {
	SetDutyCycle(percentage float64)
	GetDutyCycle() float64
}

// Instance implements the Pwm interface
type Instance struct {
	pin         *sysfs.PWMPin
	frequencyHz float64
}

func (p *Instance) getPwmPeriodNanoseconds() uint32 {
	periodInNanoseconds := (1000000000.0 / p.frequencyHz)
	return uint32(periodInNanoseconds)
}

// SetDutyCycle sets the pwm values for duty cycle, percentage is a value from 0 to 100.0
func (p *Instance) SetDutyCycle(percentage float64) error {
	dutyCycleInNanoseconds := (percentage * float64(p.getPwmPeriodNanoseconds())) / 100.0
	err := p.pin.SetDutyCycle(uint32(dutyCycleInNanoseconds))

	return err
}

// GetDutyCycle retrieves the current pwm values and returns a value from 0 to 100.0
func (p *Instance) GetDutyCycle() (float64, error) {
	periodNanoSeconds, err := p.pin.DutyCycle()
	if err != nil {
		return 0, err
	}

	// convert nanoseconds to duty cycle
	dutyCycle := (float64(periodNanoSeconds) / float64(p.getPwmPeriodNanoseconds())) * 100.0
	return dutyCycle, nil
}

// New creates a new pwm instance
func New(logger *zap.SugaredLogger, channel int, frequencyHz float64) (*Instance, error) {
	pin := sysfs.NewPWMPin(channel)

	i := &Instance{
		pin:         pin,
		frequencyHz: frequencyHz,
	}

	err := pin.Export()
	if err != nil {
		logger.Errorw(err.Error(), "pwm pin export failed", channel)
		return nil, err
	}

	// ensure that duty_cycle is lower than period, see package notes
	initialDutyCycle := uint32(0)
	err = pin.SetDutyCycle(initialDutyCycle)
	if err != nil {
		logger.Errorw(err.Error(), "pwm initial duty_cycle failed", initialDutyCycle)
		return nil, err
	}

	periodNanoseconds := i.getPwmPeriodNanoseconds()
	err = pin.SetPeriod(periodNanoseconds)
	if err != nil {
		logger.Infow(err.Error(), "set period", periodNanoseconds)
		return nil, err
	}

	err = pin.Enable(true)
	if err != nil {
		logger.Infow(err.Error(), "enable", "failed")
		return nil, err
	}

	return i, nil
}
