package srtm

import (
	"image"
	"testing"
)

var emptySRTM1Img *SRTMImage = &SRTMImage{Format: SRTM1Format, Data: make([]int16, SRTM1Size*SRTM1Size)}
var emptySRTM3Img *SRTMImage = &SRTMImage{Format: SRTM3Format, Data: make([]int16, SRTM3Size*SRTM3Size)}

func TestIndexToCoordinatesSRTM1(t *testing.T) {
	var point image.Point
	point = emptySRTM1Img.IndexToCoordinates(0)
	if point.X != 0 || point.Y != 0 {
		t.Error("IndexToCoordinates(0) should return (0,0), but returned", point.X, point.Y)
	}
	point = emptySRTM1Img.IndexToCoordinates(1)
	if point.X != 1 || point.Y != 0 {
		t.Error("IndexToCoordinates(1) should return (1,0), but returned", point.X, point.Y)
	}
	point = emptySRTM1Img.IndexToCoordinates(3600)
	if point.X != 3600 || point.Y != 0 {
		t.Error("IndexToCoordinates(3600) should return (3600,0), but returned", point.X, point.Y)
	}
	// Note we are 0-indexed, so the last index is 3600, not 3601, which is in the next row.
	point = emptySRTM1Img.IndexToCoordinates(3601)
	if point.X != 0 || point.Y != 1 {
		t.Error("IndexToCoordinates(3601) should return (0,1), but returned", point.X, point.Y)
	}
	point = emptySRTM1Img.IndexToCoordinates(7201)
	if point.X != 3600 || point.Y != 1 {
		t.Error("IndexToCoordinates(7201) should return (3600,1), but returned", point.X, point.Y)
	}
	point = emptySRTM1Img.IndexToCoordinates(12967200)
	if point.X != 3600 || point.Y != 3600 {
		t.Error("IndexToCoordinates(25934402) should return (3600,3600), but returned", point.X, point.Y)
	}
}

func TestIndexToCoordinateSRTM3(t *testing.T) {
	var point image.Point
	point = emptySRTM3Img.IndexToCoordinates(0)
	if point.X != 0 || point.Y != 0 {
		t.Error("IndexToCoordinates(0) should return (0,0), but returned", point.X, point.Y)
	}
	point = emptySRTM3Img.IndexToCoordinates(1)
	if point.X != 1 || point.Y != 0 {
		t.Error("IndexToCoordinates(1) should return (1,0), but returned", point.X, point.Y)
	}
	point = emptySRTM3Img.IndexToCoordinates(1200)
	if point.X != 1200 || point.Y != 0 {
		t.Error("IndexToCoordinates(1200) should return (1200,0), but returned", point.X, point.Y)
	}
	// Note we are 0-indexed, so the last index is 1200, not 1201, which is in the next row.
	point = emptySRTM3Img.IndexToCoordinates(1201)
	if point.X != 0 || point.Y != 1 {
		t.Error("IndexToCoordinates(1201) should return (0,1), but returned", point.X, point.Y)
	}
	point = emptySRTM3Img.IndexToCoordinates(2401)
	if point.X != 1200 || point.Y != 1 {
		t.Error("IndexToCoordinates(2401) should return (1200,1), but returned", point.X, point.Y)
	}
	point = emptySRTM3Img.IndexToCoordinates(1442400)
	if point.X != 1200 || point.Y != 1200 {
		t.Error("IndexToCoordinates(432400) should return (1200,1200), but returned", point.X, point.Y)
	}
}
