package main

import (
	"encoding/hex"
	"flag"
	"github.com/leoluk/paperlike-go/pkg/dasung"
	"log"
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

	for _, v := range flag.Args() {
		b, err := hex.DecodeString(v)
		if err != nil {
			log.Fatal(err)
		}

		if err := d.RawSetVCP(b); err != nil {
			log.Print(err)
		}
	}
}
