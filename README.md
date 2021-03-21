# paperlike-go

paperlike-go is a Go library and CLI tool to control a Dasung Paperlike HD* (2019) screen.

Installation instructions:

    git clone github.com/leoluk/paperlike-go && cd paperlike-go
    go build github.com/leoluk/paperlike-go/cmd/paperlike-cli

    sudo install -m 0755 paperlike-cli /usr/local/bin/paperlike-cli


Your local user must be able to write to the monitor's i2c bus. First, figure
out the right bus using third-party ddcutil tool (likely packaged by your distro):

    ddcutil detect --verbose

Create a udev rule which grants your user's group write permissions to i2c devices:

    echo "SUBSYSTEM==\"i2c-dev\", GROUP=\"your-group-here\", MODE=\"0660\"" | sudo tee /etc/udev/rules.d/50-i2c.rules
    udevadm trigger
    
Usage:

    Usage of paperlike-cli:
      -clear
            Clear screen
      -contrast int
            Set contrast (1-9)
      -i2c string
            i2c device path (see ddcutil detect --verbose)
      -mode int
            Set dithering mode (1-4)
      -speed int
            Set drawing speed (1-5)

Dithering modes:

    1	Fast++
    2	Fast+
    3	Fast
    4	Black+
    5	Black++

Example i3 config:

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
