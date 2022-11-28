package srtm

import (
	"image"
)

// DataVoidIndices returns points of all voids in the srtmImage.
// Data voids are represented by the value -32768 as per the SRTM documentation.
func (s *SRTMImage) DataVoidPoints() []image.Point {
	var points []image.Point
	for i, v := range s.Data {
		if v == -32768 {
			points = append(points, s.IndexToCoordinates(i))
		}
	}
	return points
}

// MinMaxElevation returns the minimum and maximum elevation values.
// Data voids are ignored and not interpreted as minimum.
// Values may be erroneous, because of other invalid data.
func (s *SRTMImage) MinMaxElevation() (min int16, max int16) {
	// do not forget to initialize min and max
	min = 32767
	max = -32768

	for _, v := range s.Data {
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

// MeanElevation returns the mean elevation value.
// Overflows as well as voids are mitigated.
// Thus may not be the actual mean.
func (s *SRTMImage) MeanElevation() int16 {
	var avg = 0

	for i, v := range s.Data {
		// avoid adding voids, by just adding the average
		if v == -32768 {
			v = int16(avg)
		}
		avg += (int(v) - avg) / (i + 1)
	}
	return int16(avg)
}

// IndexToCoordinates converts an index into the data array of the given SRTMImage
// to x,y coordinates.
func (srtmImg *SRTMImage) IndexToCoordinates(index int) image.Point {
	return IndexToCoordinates(index, srtmImg.Format)
}

// IndexToCoordinates converts an index into the data array of a SRTMImage
// with the given format to x,y coordinates.
func IndexToCoordinates(index int, format SRTMFormat) image.Point {
	x := index % format.Size()
	y := index / format.Size()
	return image.Point{x, y}
}

// CoordinatesToIndex converts x,y coordinates to an index into the data array
// for the given SRTMImage.
func (srtmImg *SRTMImage) CoordinatesToIndex(point image.Point) int {
	return CoordinatesToIndex(point, srtmImg.Format)
}

// CoordinatesToIndex converts x,y coordinates to an index into the data array
// for the given SRTM format.
func CoordinatesToIndex(point image.Point, format SRTMFormat) int {
	return point.Y*format.Size() + point.X
}

// ElevationAt returns the elevation value at the given coordinates.
func (s *SRTMImage) ElevationAt(x, y int) int16 {
	return s.Data[s.CoordinatesToIndex(image.Point{x, y})]
}
