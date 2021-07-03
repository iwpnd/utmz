package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	utmz "github.com/iwpnd/utmz"
)

func main() {

	output := os.Stderr
	flag.Usage = func() {
		fmt.Fprintf(output, `

Usage: utmz [--lat latitude] [--lng longitude]

Basic Options:
--lat latitude : Latitude of Point to get EPSG code for
--lng longitude : Longitude of Point to get EPSG code for
`)
	}

	var lat float64
	var lng float64

	flag.Float64Var(&lat, "lat", -9999, "Latitude")
	flag.Float64Var(&lng, "lng", -9999, "Longitude")

	flag.Parse()

	p := utmz.Point{Lng: lng, Lat: lat}

	if !p.Valid() {
		err := errors.New("Invalid point")
		fmt.Println(err)
		return
	}

	fmt.Printf("EPSG:%v\n", utmz.Epsg(p))
}
