package utmz

import (
	"math"
)

// Point ...
type Point struct {
	Lng, Lat float64
}

// Valid ...
func (p Point) Valid() bool {
	return p.Lng >= -180 && p.Lng <= 180 && p.Lat >= -90 && p.Lat <= 90
}

// Zone ...
func Zone(p Point) int {

	if p.Lat >= 72.0 && p.Lat < 84.0 {
		if p.Lng >= 0.0 && p.Lng < 9.0 {
			return 31
		}
		if p.Lng >= 9.0 && p.Lng < 21.0 {
			return 33
		}
		if p.Lng >= 21.0 && p.Lng < 33.0 {
			return 35
		}
		if p.Lng >= 33.0 && p.Lng < 42.0 {
			return 37
		}
	}

	if p.Lat >= 56.0 && p.Lat < 64.0 && p.Lng >= 3.0 && p.Lng <= 12.0 {
		return 32
	}

	return int(math.Round((183 + p.Lng) / 6))
}

// Epsg ...
func Epsg(p Point) int {
	zone := Zone(p)

	if p.Lat > 0 {
		return 32600 + zone
	}

	return 32700 + zone
}
