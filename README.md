# Overview

Example golang application that makes use of the hardware PWM channels.

Written as an example for the Onion Omega2 but this would serve as a good example for a number of
compute modules that have PWM hardware exposed to Linux.

# Dependencies

gobot.io's sysfs module is used to provide access to the pwm hardware

# How to build

The project has been updated to take advantage of go's new module support. This makes it
easy to build.

```
<go to where you'd like to place the 'go-pwm-example' directory>
git clone https://github.com/chmorgan/go-pwm-example.git
cd go-pwm-example
go build
```

## Build for the Omega2

golang has support for a number of architectures built in, this makes it well suited for cross platform
development.

The Omega2 uses a MediaTek MT7688 SoC. This processor is a mips architecture and little endian is being used.
Therefore the architecture is 'mips' + 'le', or 'mipsle'.

```
GOOS=linux GOARCH=mipsle go build
```

You should now have a 'go-pwm-example' application in the current directory.

## Copy to your Omega2

Replace 'xxxx' with the address of your Omega.

```
scp go-pwm-example root@omega-xxxx.local:/root/
```

# How to run

## Pin mux configuration
Note: Pin muxing must be set appropriately for the pwm channel being used.

For example:
```
# omega2-ctrl gpiomux get
Group i2c - [i2c] gpio
Group uart0 - [uart] gpio
Group uart1 - [uart] gpio pwm01
Group uart2 - [uart] gpio pwm23
Group pwm0 - pwm [gpio]
Group pwm1 - pwm [gpio]
Group refclk - refclk [gpio]
Group spi_s - spi_s [gpio] pwm01_uart2
Group spi_cs1 - [spi_cs1] gpio refclk
Group i2s - i2s [gpio] pcm
Group ephy - [ephy] gpio
Group wled - wled [gpio]
```

Assuming we are using pwm channel 0 (GPIO18) you'll reconfigure this pin for pwm mode via:
```
# omega2-ctrl gpiomux set pwm0 pwm
set gpiomux pwm0 -> pwm
```

## Running
```
./go-pwm-example
```

And to get command line help you can pass the '-help' option:

```
# ./go-pwm-example -help
Usage of ./go-pwm-example:
  -channel int
    	PWM channel (0, 1, 2, or 3)
  -frequency float
    	Frequency in Hz (default 100)
```
