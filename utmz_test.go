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
		got, _ := Epsg(test.lat, test.lng)

		if got != test.expected {
			t.Errorf("Expected %v, got: %v", test.expected, got)
		}
	}
}

func TestEpsgErr(t *testing.T) {
	got, _ := Epsg(91, 181)

	if got != 0 {
		t.Errorf("Expected error, got %v", got)
	}
}

func TestInvalidPointError(t *testing.T) {
	expected := "Point{Lat: 52, Lng: -190} - invalid point"
	_, err := Epsg(52, -190)

	if err.Error() != expected {
		t.Errorf("Expected: %v, got: %v", expected, err)

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

func TestProj4(t *testing.T) {
	tests := []struct {
		lat, lng float64
		expected string
	}{
		{lat: 52, lng: 13, expected: "+proj=utm +zone=33 +datum=WGS84 +units=m +no_def +ellps=WGS84 +towgs84=0,0,0"},
		{lat: -52, lng: 13, expected: "+proj=utm +zone=33S +datum=WGS84 +units=m +no_def +ellps=WGS84 +towgs84=0,0,0"},
	}

	for _, test := range tests {
		got, _ := Proj4(test.lat, test.lng)

		if got != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, got)
		}
	}
}
