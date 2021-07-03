package utmz

import "testing"

func TestEpsg(t *testing.T) {
	if Epsg(Point{Lat: 52, Lng: 13}) != 32633 {
		t.Error("Expected EPSG:32633")
	}
}
