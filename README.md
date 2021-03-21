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
