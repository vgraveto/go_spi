package main

import (
	"golang.org/x/exp/io/spi"
	"fmt"
	"time"
)

func main() {
	dev, err := spi.Open(&spi.Devfs{
		Dev:      "/dev/spidev1.0",
		Mode:     spi.Mode0,
		MaxSpeed: 3600000,
	})
	if err != nil {
		panic(err)
	}
	defer dev.Close()


	txbuf := []byte{0x01, 0x80, 0x00}
	rxbuf := make([]byte, 3)

	var (
		value uint16
		avg float64
		)
	max := 10000;
	for i:=1;i<=10;i++ {
		iniT := time.Now()
		for n:=0; n<max*i; n++ {
			if err := dev.Tx(txbuf, rxbuf); err != nil {
				panic(err)
			}
			value = uint16(rxbuf[1] & 0x03) << 8 | uint16(rxbuf[2])
			//		fmt.Printf("%d\n", value)
		}
		fimT := time.Now()
		difT := fimT.Sub(iniT)
		samplesPerSec := 1E9*float64(max*i)/float64(difT)
		avg += samplesPerSec
		fmt.Printf("Duration: %v - %d samples - %f samples/s\n", difT, max*i, samplesPerSec)
	}

	fmt.Printf("Avg = %f samples/s\n",avg/10)
	fmt.Printf("Last value=%d\n",value)
}
