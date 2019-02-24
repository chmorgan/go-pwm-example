# Overview

Example golang application that makes use of the hardware PWM channels.

Written as an example for the Onion Omega2 but this would serve as a good example for a number of
compute modules that have PWM hardware exposed to Linux.

# Dependencies

gobot.io's sysfs module is used to provide access to the pwm hardware

# How to build (for inexperienced golang users)

golang has particular conventions for its approach to workspaces.

* Determine where your golang workspace will be, a good location is ~/go. The rest of these steps assume
that directory is your workspace. Swap it out with your actual workspace directory if it differs.

```
mkdir ~/go
cd ~/go
mkdir src
export GOPATH=`pwd`
go get https://github.com/chmorgan/go-pwm-example
cd src/github.com/chmorgan/go-pwm-example
go get .
```

## Build for the Omega2

golang has support for a number of architectures built in, this makes it well suited for cross platform
development.

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
