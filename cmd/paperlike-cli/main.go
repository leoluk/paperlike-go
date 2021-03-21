package main

import (
	"flag"
	"github.com/leoluk/paperlike-go/pkg/dasung"
	"log"
	"path/filepath"
	"strings"
)

var (
	flagDevicePath = flag.String("i2c", "", "i2c device path (see ddcutil detect --verbose)")

	flagContrast = flag.Int("contrast", 0, "Set contrast (1-9)")
	flagMode     = flag.Int("mode", 0, "Set dithering mode (1-4)")
	flagSpeed    = flag.Int("speed", 0, "Set drawing speed (1-5)")

	flagClear = flag.Bool("clear", false, "Clear screen")
)

func init() {
	flag.Parse()

	if *flagDevicePath == "" {
		log.Fatal("please specify device path")
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

	if *flagClear {
		if err := d.ClearScreen(); err != nil {
			log.Fatal(err)
		}
	}
}
