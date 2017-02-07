package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num float64
	var decimal float64
	var dms float64

	num, err := strconv.ParseFloat(os.Args[1], 64) // Convert string to float64
	if err == nil {
		decimal = dms2dec(num)
		fmt.Printf("Decimal is: %.9f\n", decimal)
		dms = dec2dms(decimal)
		fmt.Printf("Degrees Minutes Seconds is: %.6f\n", dms)
	}

}

// Convert dd.mmss to decimal degrees
func dms2dec(dms float64) float64 {

	sign := 1.0
	if dms < 0 {
		sign = -1.0
		dms = dms * sign
	}

	degrees := float64(int(dms))
	minutes := float64(int((dms * 100))) - degrees*100
	seconds := (((dms - degrees) * 100) - minutes) * 100
	decdeg := degrees + minutes/60 + seconds/3600
	return decdeg * sign
}

// Convert decimal degrees to dd.mmss (degrees,minutes,seconds)
func dec2dms(dec float64) float64 {

	sign := 1.0
	if dec < 0 {
		sign = -1.0
		dec = dec * sign
	}

	degrees := float64(int(dec))
	minutes := float64(int((dec * 60))) - degrees*60
	seconds := ((dec - degrees) * 3600) - minutes*60
	dms := degrees + minutes/100 + seconds/10000
	return dms * sign
}
