package main

import (
	"flag"
	"github.com/jtbg/paperlike-go/dasung"
	"log"
	"path/filepath"
	"strings"
)

var (
	flagDevicePath = flag.String("i2c", "", "i2c device path (see ddcutil detect --verbose)")

	flagContrast    = flag.Int("contrast", 0, "Set contrast (1-9)")
	flagMode        = flag.Int("mode", 0, "Set dithering mode (1-4)")
	flagSpeed       = flag.Int("speed", 0, "Set drawing speed (1-5)")
	light1Intensity = flag.Int("light1", -1, "Set light1 intensity (0-85)")
	light2Intensity = flag.Int("light2", -1, "Set light2 intensity (0-85)")

	flagClear = flag.Bool("clear", false, "Clear screen")
)

func init() {
	flag.Parse()

	if *flagDevicePath == "" {
		path, err := dasung.FindDasungI2CDevicePath()
		if err != nil {
			log.Fatal("Failed to find Dasung Paperlike display:", err)
		}
		*flagDevicePath = path
	}

	if !strings.HasPrefix(filepath.Clean(*flagDevicePath), "/dev/i2c-") {
		log.Fatal("invalid device path (must start with /dev/i2c-)")
	}
}

func main() {
	d, err := dasung.NewDasungControl(*flagDevicePath)
	if err != nil {
		log.Fatal(err)
	}

	if *flagContrast != 0 {
		if err := d.SetContrast(*flagContrast); err != nil {
			log.Fatal(err)
		}
	}

	if *flagMode != 0 {
		if err := d.SetDitheringMode(dasung.DitheringMode(*flagMode)); err != nil {
			log.Fatal(err)
		}
	}

	if *flagSpeed != 0 {
		if err := d.SetRefreshSpeed(dasung.RefreshSpeed(*flagSpeed)); err != nil {
			log.Fatal(err)
		}
	}

	if *light1Intensity > -1 {
		if err := d.SetLightIntensity(1, *light1Intensity); err != nil {
			log.Fatal(err)
		}
	}

	if *light2Intensity > -1 {
		if err := d.SetLightIntensity(2, *light2Intensity); err != nil {
			log.Fatal(err)
		}
	}

	if *flagClear {
		if err := d.ClearScreen(); err != nil {
			log.Fatal(err)
		}
	}
}
