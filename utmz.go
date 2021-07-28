package utmz

import (
	"fmt"
	"math"
)

type invalidPointError struct {
	p   Point
	msg string
}

func (e *invalidPointError) Error() string {
	return fmt.Sprintf("Point{Lat: %v, Lng: %v} - %s", e.p.Lat, e.p.Lng, e.msg)
}

// Point ...
type Point struct {
	Lng, Lat float64
}

// Valid to validate a point
func (p Point) Valid() bool {
	return p.Lng >= -180 && p.Lng <= 180 && p.Lat >= -90 && p.Lat <= 90
}

// Zone to get UTM zone from point
func Zone(p Point) (int, error) {

	if !p.Valid() {
		return 0, &invalidPointError{p: p, msg: "invalid point"}
	}

	if p.Lat >= 72.0 && p.Lat < 84.0 {
		if p.Lng >= 0.0 && p.Lng < 9.0 {
			return 31, nil
		}
		if p.Lng >= 9.0 && p.Lng < 21.0 {
			return 33, nil
		}
		if p.Lng >= 21.0 && p.Lng < 33.0 {
			return 35, nil
		}
		if p.Lng >= 33.0 && p.Lng < 42.0 {
			return 37, nil
		}
	}

	if p.Lat >= 56.0 && p.Lat < 64.0 && p.Lng >= 3.0 && p.Lng <= 12.0 {
		return 32, nil
	}

	return int(math.Round((183 + p.Lng) / 6)), nil
}

// Epsg of utm zone from point
func Epsg(p Point) (int, error) {
	zone, err := Zone(p)

	if err != nil {
		return 0, err
	}

	if p.Lat > 0 {
		return 32600 + zone, nil
	}

	return 32700 + zone, nil
}

// Proj4 to get proj4 format from point
func Proj4(p Point) (string, error) {
	var zs string

	z, err := Zone(p)

	if err != nil {
		return "", err
	}

	// if souther hemisphere add an S to zone number
	// else leave blank zone number defaulting to north
	if p.Lat < 0 {
		zs = fmt.Sprintf("%vS", z)
	} else {
		zs = fmt.Sprint(z)
	}

	return fmt.Sprintf("+proj=utm +zone=%v +datum=WGS84 +units=m +no_def +ellps=WGS84 +towgs84=0,0,0", zs), nil
}
