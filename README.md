# paperlike-go

[![GoDoc](https://godoc.org/github.com/leoluk/paperlike-go?status.svg)](https://godoc.org/github.com/leoluk/paperlike-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/leoluk/paperlike-go)](https://goreportcard.com/report/github.com/leoluk/paperlike-go)



paperlike-go is a Go library and CLI tool to control a Dasung Paperlike HD* (2019) screen.

# Installation
Paperlike-go can be installed using the `go install` command:

    sudo GOBIN=/usr/local/bin/ go install github.com/leoluk/paperlike-go/cmd/paperlike-cli

## Set up i2c permission for your user
Your local user must be able to write to the monitor's i2c bus. 

1. Create a new user group called `i2c`:

       sudo groupadd i2c

2. Add your user to the `i2c` group:

       sudo usermod -a -G i2c $USER

3. Change the group ownership of `/dev/i2c` codepoints to the `i2c` group:

       sudo chown :i2c /dev/i2c-* 
    
4. Update the group permissions of those same to allow read and write access:

       sudo chmod g+rw /dev/i2c-*

5. Create a udev rule to make these changes persistent:

       sudo echo 'KERNEL=="i2c-[0-9]*", GROUP="i2c"' >> /etc/udev/rules.d/10-local_i2c_group.rules

## Find the i2c path for your Dasung device

To determine the i2c path for your device, run the following command:

    ddcutil detect --brief | grep -B 1 -Ei "dsc|dasung|paper"

Make note of the I2C bus path (e.g. `/dev/i2c-3`).

>In the case that that this command gives an ambiguous or empty result, you can read the full output by running `ddcutil detect --verbose`

# Usage:

    Usage: paperlike-cli -i2c <i2c path> [command<value>] 

      Mandatory flags:  
            -i2c <path>       absolute path of the target i2c device

      Commands:
            -clear            Clear & refresh screen
            -contrast <1-9>   Set contrast
            -light1 <0-85>    Set light1 intensity; -1 to disable (default)
            -light2 <0-85>    Set light2 intensity; -1 to disable (default)
            -mode <1-4>       Set dithering/display mode
            -speed <1-5>      Set drawing speed

      Examples:
            paperlike-cli -i2c /dev/i2c-3 -mode 1     set the display mode to M1 for the device at /dev/i2c-3
            paperlike-cli -i2c /dev/i2c-5 -clear      clear/refresh the screen for the device at /dev/i2c-5

## Display modes:

>Refer to your device's manual for the mapping of M1-M4 to specific image modes.

<details>
      <summary>
            Dasung Paperlike HD (gen 3) Modes
      </summary>

| Mode | Name | Goal |Usecases| Notes |
|:-:|:--|:-- |:--| :--|
|1|A2|Emphasize lower latency|Chat applications, real-time dashboards|If ghosting is an issue, make use of the `-clear` command|
|2| [Floyd](https://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering)| Balance latency and detail |Web browsing, coding, text composition| For best performance, set contrast 3-6.
|3|A16|Emphasize fine details|Reading PDFs, calendars, task lists|The Accessibility "Reduce Animations" options in your OS may be useful|
|4|-|-| 

</details>
<details>
      <summary>
            Dasung Paperlike HD-F/HD-FT Modes
      </summary>

| Mode | Name | Goal |Usecases| Notes |
|:-:|:--|:-- |:--| :--|
|1|A2|Emphasize lower latency|Chat applications, real-time dashboards|If ghosting is an issue, make use of the `-clear` command|
|2| [Floyd](https://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering)| Balance latency and detail |Web browsing, coding, text composition| For best performance, set contrast 3-6.
|3|A61|Emphasize fine details|Images, PDFs, handwriting|The Accessibility "Reduce Animations" options in your OS may be useful|
|4|-|-| 

</details>

## Drawing speeds:

    1	Fast++
    2	Fast+
    3	Fast
    4	Black+
    5	Black++

## Example i3 config:

    bindsym $mod+Mod1+1 exec paperlike-cli -i2c /dev/i2c-0 -contrast 1
    bindsym $mod+Mod1+2 exec paperlike-cli -i2c /dev/i2c-0 -contrast 2
    bindsym $mod+Mod1+3 exec paperlike-cli -i2c /dev/i2c-0 -contrast 3
    bindsym $mod+Mod1+4 exec paperlike-cli -i2c /dev/i2c-0 -contrast 4
    bindsym $mod+Mod1+5 exec paperlike-cli -i2c /dev/i2c-0 -contrast 5
    bindsym $mod+Mod1+6 exec paperlike-cli -i2c /dev/i2c-0 -contrast 6
    bindsym $mod+Mod1+7 exec paperlike-cli -i2c /dev/i2c-0 -contrast 7
    bindsym $mod+Mod1+8 exec paperlike-cli -i2c /dev/i2c-0 -contrast 8
    bindsym $mod+Mod1+9 exec paperlike-cli -i2c /dev/i2c-0 -contrast 9
    bindsym $mod+Mod1+h exec paperlike-cli -i2c /dev/i2c-0 -mode 1
    bindsym $mod+Mod1+j exec paperlike-cli -i2c /dev/i2c-0 -mode 2
    bindsym $mod+Mod1+k exec paperlike-cli -i2c /dev/i2c-0 -mode 3
    bindsym $mod+Mod1+l exec paperlike-cli -i2c /dev/i2c-0 -mode 4
    bindsym $mod+Mod1+z exec paperlike-cli -i2c /dev/i2c-0 -speed 1
    bindsym $mod+Mod1+u exec paperlike-cli -i2c /dev/i2c-0 -speed 2
    bindsym $mod+Mod1+i exec paperlike-cli -i2c /dev/i2c-0 -speed 3
    bindsym $mod+Mod1+o exec paperlike-cli -i2c /dev/i2c-0 -speed 4
    bindsym $mod+Mod1+p exec paperlike-cli -i2c /dev/i2c-0 -speed 5
    bindsym $mod+Mod1+Return exec paperlike-cli -i2c /dev/i2c-0 -clear
