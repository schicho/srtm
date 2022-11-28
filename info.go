package srtm

import (
	"errors"
	"fmt"
	"image"
	"sort"
)

var ErrPointOutOfBounds = errors.New("point out of bounds for SRTM image format")
var ErrIndexOutOfBounds = errors.New("index out of bounds for SRTM image format")

// DataVoidIndices returns points of all voids in the srtmImage.
// Data voids are represented by the value -32768 as per the SRTM documentation.
func (srtmImg *SRTMImage) ElevationVoids() []image.Point {
	var points []image.Point
	for i, v := range srtmImg.Data {
		if v == -32768 {
			point, _ := IndexToCoordinates(i, srtmImg.Format)
			points = append(points, point)
		}
	}
	return points
}

// ElevationMinMax returns the minimum and maximum elevation values.
// Data voids are ignored and not interpreted as minimum.
// Values may be erroneous, because of other invalid data.
func (srtmImg *SRTMImage) ElevationMinMax() (min int16, max int16) {
	// do not forget to initialize min and max
	min = 32767
	max = -32768

	for _, v := range srtmImg.Data {
		// avoid letting voids influence the min/max
		if v < min && v != -32768 {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

// ElevationMean returns the mean elevation value.
// Overflows as well as voids are mitigated.
// Thus may not be the actual mean.
func (srtmImg *SRTMImage) ElevationMean() int16 {
	var avg = 0

	for i, v := range srtmImg.Data {
		// avoid adding voids, by just adding the average
		if v == -32768 {
			v = int16(avg)
		}
		avg += (int(v) - avg) / (i + 1)
	}
	return int16(avg)
}

// ElevationAt returns the elevation value at the given coordinates.
func (srtmImg *SRTMImage) ElevationAt(point image.Point) (int16, error) {
	index, err := CoordinatesToIndex(point, srtmImg.Format)
	if err != nil {
		return -1, err
	}
	return srtmImg.Data[index], nil
}

// ElevationPercentile returns the nearest-ranked percentile of the elevation values.
// Percentile must be between 0 and 1.
// Data voids are not ignored and part of the percentile calculation.
func (srtmImg *SRTMImage) ElevationPercentile(percentile float64) int16 {
	if percentile < 0 || percentile > 1 {
		panic(fmt.Sprintf("percentile must be between 0 and 1, but was %f", percentile))
	}
	data := make([]int16, len(srtmImg.Data))
	copy(data, srtmImg.Data)
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })

	if percentile == 1 {
		return data[len(data)-1]
	}

	index := int(float64(len(data)) * percentile)
	return data[index]
}

// IndexToCoordinates converts an index into the data array of a SRTMImage
// with the given format to x,y coordinates.
func IndexToCoordinates(index int, format SRTMFormat) (image.Point, error) {
	if !IsIndexInBounds(index, format) {
		return image.Point{}, fmt.Errorf("%w: %v, index: %v", ErrIndexOutOfBounds, format, index)
	}
	x := index % format.Size()
	y := index / format.Size()
	return image.Point{x, y}, nil
}

// CoordinatesToIndex converts x,y coordinates to an index into the data array
// for the given SRTMFormat.
func CoordinatesToIndex(point image.Point, format SRTMFormat) (int, error) {
	if !IsPointInBounds(point, format) {
		return -1, fmt.Errorf("%w: %v, point: %v", ErrPointOutOfBounds, format, point)
	}
	return point.Y*format.Size() + point.X, nil
}

// IsPointInBounds checks if the given point is inside the data array of the given SRTMFormat.
func IsPointInBounds(point image.Point, format SRTMFormat) bool {
	return point.X >= 0 && point.X < format.Size() && point.Y >= 0 && point.Y < format.Size()
}

// IsIndexInBounds returns true if the given index is inside the data array of the given SRTMFormat.
func IsIndexInBounds(index int, format SRTMFormat) bool {
	return index >= 0 && index < format.Size()*format.Size()
}
