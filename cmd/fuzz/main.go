package main

import (
	"encoding/hex"
	"flag"
	"github.com/leoluk/paperlike-go/pkg/dasung"
	"log"
	"time"
)

var (
	flagDevicePath = flag.String("i2c", "", "i2c device path (see ddcutil detect --verbose)")
)

func init() {
	flag.Parse()

	if *flagDevicePath == "" {
		log.Fatal("please specify device path")
	}
}

func main() {
	d, err := dasung.NewDasungControl(*flagDevicePath)
	if err != nil {
		log.Fatal(err)
	}

	for i := uint8(0); i < 0x03; i++ {
		for j := uint8(0); j < 0x03; j++ {
			for k := uint8(0); k < 0x03; k++ {
				for l := uint8(0); l < 0x03; l++ {
					b := []byte{i, j, k, l}

					if k == 0x0c { // skip clear
						continue
					}

					log.Print(hex.EncodeToString(b))

					if err := d.RawSetVCP(b); err != nil {
						log.Print(err)
					}

					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}
}
