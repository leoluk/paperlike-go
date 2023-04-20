package dasung

import (
	"bytes"
	"errors"
	"golang.org/x/exp/io/i2c"
	"log"
)

const (
	// Well-known DDC-CI I2C bus address.
	i2cDDCAddr = 0x37
)

type DasungControl struct {
	d *i2c.Device
}

// NewDasungControl opens the I2C device.
func NewDasungControl(device string) (*DasungControl, error) {
	d, err := i2c.Open(&i2c.Devfs{Dev: device}, i2cDDCAddr)
	if err != nil {
		return nil, err
	}

	return &DasungControl{d}, nil
}

func (d *DasungControl) RawSetVCP(arg []byte) error {
	// Set VCP Feature (0x03), Feature: 0x08, two filler bytes
	b := bytes.NewBuffer([]byte{0x03, 0x08, 0x00, 0x00})

	if _, err := b.Write(arg); err != nil {
		panic(err)
	}

	log.Printf("sending: %x", b)

	if err := d.d.Write(b.Bytes()); err != nil {
		return err
	}

	r := make([]byte, 11)
	if err := d.d.Read(r); err != nil {
		return err
	}

	log.Printf("received: %x (len %d)", r, len(r))

	return nil
}

type DitheringMode int

const (
	ModeM1 DitheringMode = iota + 1
	ModeM2
	ModeM3
	ModeM4
)

// SetDitheringMode sets the display dithering mode to the specified mode (M1-M4).
func (d *DasungControl) SetDitheringMode(mode DitheringMode) error {
	if mode < 1 || mode > 4 {
		return errors.New("mode out of range")
	}

	return d.RawSetVCP([]byte{0x07, uint8(mode)})
}

// SetContrast sets the contrast to a value between 1 and 9.
func (d *DasungControl) SetContrast(level int) error {
	if level < 1 || level > 9 {
		return errors.New("level out of range")
	}

	return d.RawSetVCP([]byte{0x08, uint8(level)})
}

type RefreshSpeed int

const (
	SpeedFastPlusPlus RefreshSpeed = iota + 1
	SpeedFastPlus
	SpeedFast
	SpeedBlackPlus
	SpeedBlackPlusPlus
)

// SetRefreshSpeed sets the specified refresh mode.
func (d *DasungControl) SetRefreshSpeed(mode RefreshSpeed) error {
	if mode < 1 || mode > 5 {
		return errors.New("level out of range")
	}

	return d.RawSetVCP([]byte{0x0c, uint8(mode)})
}

type LightID int

const (
	Light1 LightID = iota + 1
	Light2
)

// SetLightIntensity sets the specified light intensity
func (d *DasungControl) SetLightIntensity(light LightID, level int) error {
	if light < 1 || light > 2 {
		return errors.New("light id out of range")
	}

	if level < 0 || level > 85 {
		return errors.New("level out of range")
	}

	return d.RawSetVCP([]byte{byte(0xD + light), byte(level * 3)})
}

func (d *DasungControl) ClearScreen() error {
	return d.RawSetVCP([]byte{0x06, 0x03})
}
