package srtm

import (
	"errors"
	"image"
	"testing"
)

func TestIndexToCoordinatesSRTM1(t *testing.T) {
	var point image.Point
	point, _ = IndexToCoordinates(0, SRTM1Format)
	if point.X != 0 || point.Y != 0 {
		t.Error("IndexToCoordinates(0) should return (0,0), but returned", point)
	}
	point, _ = IndexToCoordinates(1, SRTM1Format)
	if point.X != 1 || point.Y != 0 {
		t.Error("IndexToCoordinates(1) should return (1,0), but returned", point)
	}
	point, _ = IndexToCoordinates(3600, SRTM1Format)
	if point.X != 3600 || point.Y != 0 {
		t.Error("IndexToCoordinates(3600) should return (3600,0), but returned", point)
	}
	// Note we are 0-indexed, so the last index is 3600, not 3601, which is in the next row.
	point, _ = IndexToCoordinates(3601, SRTM1Format)
	if point.X != 0 || point.Y != 1 {
		t.Error("IndexToCoordinates(3601) should return (0,1), but returned", point)
	}
	point, _ = IndexToCoordinates(7201, SRTM1Format)
	if point.X != 3600 || point.Y != 1 {
		t.Error("IndexToCoordinates(7201) should return (3600,1), but returned", point)
	}
	point, _ = IndexToCoordinates(12967200, SRTM1Format)
	if point.X != 3600 || point.Y != 3600 {
		t.Error("IndexToCoordinates(25934402) should return (3600,3600), but returned", point)
	}
}

func TestIndexToCoordinateSRTM3(t *testing.T) {
	var point image.Point
	point, _ = IndexToCoordinates(0, SRTM3Format)
	if point.X != 0 || point.Y != 0 {
		t.Error("IndexToCoordinates(0) should return (0,0), but returned", point)
	}
	point, _ = IndexToCoordinates(1, SRTM3Format)
	if point.X != 1 || point.Y != 0 {
		t.Error("IndexToCoordinates(1) should return (1,0), but returned", point)
	}
	point, _ = IndexToCoordinates(1200, SRTM3Format)
	if point.X != 1200 || point.Y != 0 {
		t.Error("IndexToCoordinates(1200) should return (1200,0), but returned", point)
	}
	// Note we are 0-indexed, so the last index is 1200, not 1201, which is in the next row.
	point, _ = IndexToCoordinates(1201, SRTM3Format)
	if point.X != 0 || point.Y != 1 {
		t.Error("IndexToCoordinates(1201) should return (0,1), but returned", point)
	}
	point, _ = IndexToCoordinates(2401, SRTM3Format)
	if point.X != 1200 || point.Y != 1 {
		t.Error("IndexToCoordinates(2401) should return (1200,1), but returned", point)
	}
	point, _ = IndexToCoordinates(1442400, SRTM3Format)
	if point.X != 1200 || point.Y != 1200 {
		t.Error("IndexToCoordinates(432400) should return (1200,1200), but returned", point.X, point.Y)
	}
}

func TestCoordinatesToIndexSRTM1(t *testing.T) {
	var index int
	index, _ = CoordinatesToIndex(image.Point{0, 0}, SRTM1Format)
	if index != 0 {
		t.Error("CoordinatesToIndex(0,0) should return 0, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{1, 0}, SRTM1Format)
	if index != 1 {
		t.Error("CoordinatesToIndex(1,0) should return 1, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{3600, 0}, SRTM1Format)
	if index != 3600 {
		t.Error("CoordinatesToIndex(3600,0) should return 3600, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{0, 1}, SRTM1Format)
	if index != 3601 {
		t.Error("CoordinatesToIndex(0,1) should return 3601, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{3600, 1}, SRTM1Format)
	if index != 7201 {
		t.Error("CoordinatesToIndex(3600,1) should return 7201, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{3600, 3600}, SRTM1Format)
	if index != 12967200 {
		t.Error("CoordinatesToIndex(3600,3600) should return 12967200, but returned", index)
	}
}

func TestCoordinatesToIndexSRTM3(t *testing.T) {
	var index int
	index, _ = CoordinatesToIndex(image.Point{0, 0}, SRTM3Format)
	if index != 0 {
		t.Error("CoordinatesToIndex(0,0) should return 0, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{1, 0}, SRTM3Format)
	if index != 1 {
		t.Error("CoordinatesToIndex(1,0) should return 1, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{1200, 0}, SRTM3Format)
	if index != 1200 {
		t.Error("CoordinatesToIndex(1200,0) should return 1200, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{0, 1}, SRTM3Format)
	if index != 1201 {
		t.Error("CoordinatesToIndex(0,1) should return 1201, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{1200, 1}, SRTM3Format)
	if index != 2401 {
		t.Error("CoordinatesToIndex(1200,1) should return 2401, but returned", index)
	}
	index, _ = CoordinatesToIndex(image.Point{1200, 1200}, SRTM3Format)
	if index != 1442400 {
		t.Error("CoordinatesToIndex(1200,1200) should return 1442400, but returned", index)
	}
}

func TestCoordinatesToIndexSRTM1Error(t *testing.T) {
	_, err := CoordinatesToIndex(image.Point{-1, 0}, SRTM1Format)
	if err == nil {
		t.Error("CoordinatesToIndex(-1,0) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(-1,0) should return an ErrPointOufOfBounds error, but returned", err)
	}
	_, err = CoordinatesToIndex(image.Point{0, -1}, SRTM1Format)
	if err == nil {
		t.Error("CoordinatesToIndex(0,-1) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(0,-1) should return an ErrPointOufOfBounds error, but returned", err)
	}
	_, err = CoordinatesToIndex(image.Point{3601, 0}, SRTM1Format)
	if err == nil {
		t.Error("CoordinatesToIndex(3601,0) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(3601,0) should return an ErrPointOufOfBounds error, but returned", err)
	}
	_, err = CoordinatesToIndex(image.Point{0, 3601}, SRTM1Format)
	if err == nil {
		t.Error("CoordinatesToIndex(0,3601) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(0,3601) should return an ErrPointOufOfBounds error, but returned", err)
	}
}

func TestCoordinatesToIndexSRTM3Error(t *testing.T) {
	_, err := CoordinatesToIndex(image.Point{-1, 0}, SRTM3Format)
	if err == nil {
		t.Error("CoordinatesToIndex(-1,0) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(-1,0) should return an ErrPointOufOfBounds error, but returned", err)
	}
	_, err = CoordinatesToIndex(image.Point{0, -1}, SRTM3Format)
	if err == nil {
		t.Error("CoordinatesToIndex(0,-1) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(0,-1) should return an ErrPointOufOfBounds error, but returned", err)
	}
	_, err = CoordinatesToIndex(image.Point{1201, 0}, SRTM3Format)
	if err == nil {
		t.Error("CoordinatesToIndex(1201,0) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(1201,0) should return an ErrPointOufOfBounds error, but returned", err)
	}
	_, err = CoordinatesToIndex(image.Point{0, 1201}, SRTM3Format)
	if err == nil {
		t.Error("CoordinatesToIndex(0,1201) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(0,1201) should return an ErrPointOufOfBounds error, but returned", err)
	}
}

func TestIsIndexInBounds(t *testing.T) {
	if IsIndexInBounds(0, SRTM1Format) == false {
		t.Error("IsIndexInBounds(0) should return true, but returned false")
	}
	if IsIndexInBounds(12967200, SRTM1Format) == false {
		t.Error("IsIndexInBounds(12967200) should return true, but returned false")
	}
	if IsIndexInBounds(12967201, SRTM1Format) == true {
		t.Error("IsIndexInBounds(12967201) should return false, but returned true")
	}
	if IsIndexInBounds(-1, SRTM1Format) == true {
		t.Error("IsIndexInBounds(-1) should return false, but returned true")
	}

	if IsIndexInBounds(0, SRTM3Format) == false {
		t.Error("IsIndexInBounds(0) should return true, but returned false")
	}
	if IsIndexInBounds(1442400, SRTM3Format) == false {
		t.Error("IsIndexInBounds(1442400) should return true, but returned false")
	}
	if IsIndexInBounds(1442401, SRTM3Format) == true {
		t.Error("IsIndexInBounds(1442401) should return false, but returned true")
	}
	if IsIndexInBounds(-1, SRTM3Format) == true {
		t.Error("IsIndexInBounds(-1) should return false, but returned true")
	}
}

func TestIsPointInBounds(t *testing.T) {
	if IsPointInBounds(image.Point{0, 0}, SRTM1Format) == false {
		t.Error("IsPointInBounds(0,0) should return true, but returned false")
	}
	if IsPointInBounds(image.Point{3600, 3600}, SRTM1Format) == false {
		t.Error("IsPointInBounds(3600,3600) should return true, but returned false")
	}
	if IsPointInBounds(image.Point{3601, 3601}, SRTM1Format) == true {
		t.Error("IsPointInBounds(3601,3601) should return false, but returned true")
	}
	if IsPointInBounds(image.Point{-1, -1}, SRTM1Format) == true {
		t.Error("IsPointInBounds(-1,-1) should return false, but returned true")
	}

	if IsPointInBounds(image.Point{0, 0}, SRTM3Format) == false {
		t.Error("IsPointInBounds(0,0) should return true, but returned false")
	}
	if IsPointInBounds(image.Point{1200, 1200}, SRTM3Format) == false {
		t.Error("IsPointInBounds(1200,1200) should return true, but returned false")
	}
	if IsPointInBounds(image.Point{1201, 1201}, SRTM3Format) == true {
		t.Error("IsPointInBounds(1201,1201) should return false, but returned true")
	}
	if IsPointInBounds(image.Point{-1, -1}, SRTM3Format) == true {
		t.Error("IsPointInBounds(-1,-1) should return false, but returned true")
	}
}

func TestWrappedErrorMessage(t *testing.T) {
	_, err := IndexToCoordinates(-1, SRTM1Format)
	if err == nil {
		t.Error("IndexToCoordinates(-1) should return an error, but returned nil")
	}
	if errors.Is(err, ErrIndexOutOfBounds) == false {
		t.Error("IndexToCoordinates(-1) should return an ErrIndexOutOfBounds error, but returned", err)
	}
	if err.Error() != "index out of bounds for SRTM image format: SRTM1, index: -1" {
		t.Error("IndexToCoordinates(-1) should return an error with message\n"+
			"'index out of bounds for SRTM image format: SRTM1, index: -1'\n"+
			"but returned", err)
	}

	_, err = CoordinatesToIndex(image.Point{-1, -1}, SRTM1Format)
	if err == nil {
		t.Error("CoordinatesToIndex(-1,-1) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(-1,-1) should return an ErrPointOutOfBounds error, but returned", err)
	}
	if err.Error() != "point out of bounds for SRTM image format: SRTM1, point: (-1,-1)" {
		t.Error("CoordinatesToIndex(-1,-1) should return an error with message\n"+
			"'point out of bounds for SRTM image format: SRTM1, point: (-1,-1)'\n"+
			"but returned", err)
	}

	_, err = IndexToCoordinates(40000000, SRTM3Format)
	if err == nil {
		t.Error("IndexToCoordinates(40000000) should return an error, but returned nil")
	}
	if errors.Is(err, ErrIndexOutOfBounds) == false {
		t.Error("IndexToCoordinates(40000000) should return an ErrIndexOutOfBounds error, but returned", err)
	}
	if err.Error() != "index out of bounds for SRTM image format: SRTM3, index: 40000000" {
		t.Error("IndexToCoordinates(40000000) should return an error with message\n"+
			"'index out of bounds for SRTM image format: SRTM3, index: 40000000'\n"+
			"but returned", err)
	}

	_, err = CoordinatesToIndex(image.Point{4000, 4000}, SRTM3Format)
	if err == nil {
		t.Error("CoordinatesToIndex(4000,4000) should return an error, but returned nil")
	}
	if errors.Is(err, ErrPointOutOfBounds) == false {
		t.Error("CoordinatesToIndex(4000,4000) should return an ErrPointOutOfBounds error, but returned", err)
	}
	if err.Error() != "point out of bounds for SRTM image format: SRTM3, point: (4000,4000)" {
		t.Error("CoordinatesToIndex(4000,4000) should return an error with message\n"+
			"'point out of bounds for SRTM image format: SRTM3, point: (4000,4000)'\n"+
			"but returned", err)
	}
}
