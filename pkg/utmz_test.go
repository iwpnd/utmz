package utmz

import "testing"

func TestEpsg(t *testing.T) {

	var tests = []struct {
		lat      float64
		lng      float64
		expected int
	}{
		{lat: 52.25, lng: 13.37, expected: 32633},
		{lat: -52.25, lng: 13.37, expected: 32733},
		{lat: 72.11, lng: 7.1337, expected: 32631},
		{lat: 72.11, lng: 9.1337, expected: 32633},
		{lat: 72.11, lng: 21.1337, expected: 32635},
		{lat: 72.11, lng: 33.1337, expected: 32637},
		{lat: 57.11, lng: 6.1337, expected: 32632},
		{lat: 33.460, lng: -112.260, expected: 32612},
	}

	for _, test := range tests {
		got := Epsg(Point{Lat: test.lat, Lng: test.lng})

		if got != test.expected {
			t.Errorf("Expected %v, got: %v", test.expected, got)
		}
	}
}

func TestPointValid(t *testing.T) {
	tests := []struct {
		lat, lng float64
		expected bool
	}{
		{lat: 52, lng: 13, expected: true},
		{lat: 91, lng: 13, expected: false},
		{lat: 52, lng: 181, expected: false},
		{lat: -91, lng: -181, expected: false},
	}

	for _, test := range tests {
		got := Point{Lat: test.lat, Lng: test.lng}.Valid()

		if got != test.expected {
			t.Errorf("Expected %v, got: %v", test.expected, got)
		}
	}
}
