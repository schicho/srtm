package srtm

import "testing"

var emptySRTM1Img *SRTMImage = &SRTMImage{Format: SRTM1Format, Data: make([]int16, SRTM1Size * SRTM1Size)}
var emptySRTM3Img *SRTMImage = &SRTMImage{Format: SRTM3Format, Data: make([]int16, SRTM3Size * SRTM3Size)}

func TestIndexToCoordinatesSRTM1(t *testing.T){
	var x, y int
	x, y = emptySRTM1Img.IndexToCoordinates(0)
	if x != 0 || y != 0 {
		t.Error("IndexToCoordinates(0) should return (0,0), but returned", x, y)
	}
	x, y = emptySRTM1Img.IndexToCoordinates(1)
	if x != 1 || y != 0 {
		t.Error("IndexToCoordinates(1) should return (1,0), but returned", x, y)
	}
	x, y = emptySRTM1Img.IndexToCoordinates(3600)
	if x != 3600 || y != 0 {
		t.Error("IndexToCoordinates(3600) should return (3600,0), but returned", x, y)
	}
	// Note we are 0-indexed, so the last index is 3600, not 3601, which is in the next row.
	x, y = emptySRTM1Img.IndexToCoordinates(3601)
	if x != 0 || y != 1 {
		t.Error("IndexToCoordinates(3601) should return (0,1), but returned", x, y)
	}
	x, y = emptySRTM1Img.IndexToCoordinates(7201)
	if x != 3600 || y != 1 {
		t.Error("IndexToCoordinates(7201) should return (3600,1), but returned", x, y)
	}
	x, y = emptySRTM1Img.IndexToCoordinates(12967200)
	if x != 3600 || y != 3600 {
		t.Error("IndexToCoordinates(25934402) should return (3600,3600), but returned", x, y)
	}
}

func TestIndexToCoordinateSRTM3(t *testing.T) {
	var x, y int
	x, y = emptySRTM3Img.IndexToCoordinates(0)
	if x != 0 || y != 0 {
		t.Error("IndexToCoordinates(0) should return (0,0), but returned", x, y)
	}
	x, y = emptySRTM3Img.IndexToCoordinates(1)
	if x != 1 || y != 0 {
		t.Error("IndexToCoordinates(1) should return (1,0), but returned", x, y)
	}
	x, y = emptySRTM3Img.IndexToCoordinates(1200)
	if x != 1200 || y != 0 {
		t.Error("IndexToCoordinates(1200) should return (1200,0), but returned", x, y)
	}
	// Note we are 0-indexed, so the last index is 1200, not 1201, which is in the next row.
	x, y = emptySRTM3Img.IndexToCoordinates(1201)
	if x != 0 || y != 1 {
		t.Error("IndexToCoordinates(1201) should return (0,1), but returned", x, y)
	}
	x, y = emptySRTM3Img.IndexToCoordinates(2401)
	if x != 1200 || y != 1 {
		t.Error("IndexToCoordinates(2401) should return (1200,1), but returned", x, y)
	}
	x, y = emptySRTM3Img.IndexToCoordinates(1442400)
	if x != 1200 || y != 1200 {
		t.Error("IndexToCoordinates(432400) should return (1200,1200), but returned", x, y)
	}
}