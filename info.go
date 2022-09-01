package srtm

// CountDataVoids returns the number of voids in the data.
// Data voids are represented by the value -32768 as per the SRTM documentation.
func (s *SRTMImage) CountDataVoids() int {
	var count = 0
	for _, v := range s.Data {
		if v == -32768 {
			count++
		}
	}
	return count
}

// DataVoidIndices returns the indices of the voids in the data.
// Data voids are represented by the value -32768 as per the SRTM documentation.
// The indices are returned as a index to the data array. Not as x,y coordinates.
func (s *SRTMImage) DataVoidIndices() []int {
	var indices = make([]int, 0)
	for i, v := range s.Data {
		if v == -32768 {
			indices = append(indices, i)
		}
	}
	return indices
}

// IndexToCoordinates converts an index to the data array to x,y coordinates.
func (s *SRTMImage) IndexToCoordinates(index int) (x, y int) {
	x = index % s.Format.Size()
	y = index / s.Format.Size()
	return
}

// CoordinatesToIndex converts x,y coordinates to an index to the data array.
func (s *SRTMImage) CoordinatesToIndex(x, y int) int {
	return y*s.Format.Size() + x
}

// ElevationAt returns the elevation value at the given coordinates.
func (s *SRTMImage) ElevationAt(x, y int) int16 {
	return s.Data[s.CoordinatesToIndex(x, y)]
}

// MinMaxElevation returns the minimum and maximum elevation values.
// Values may be erroneous, because of voids or other invalid data.
func (s *SRTMImage) MinMaxElevation() (min int16, max int16) {	
	for _, v := range s.Data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

// MeanElevation returns the mean elevation value.
// Overflows are mitigated. As well as voids.
// Thus may not be the actual mean.
func (s *SRTMImage) MeanElevation() int16 {
	var avg = 0

	for i, v := range s.Data {
		// avoid adding voids, by just adding the average
		if v == -32768 {
			v = int16(avg)
		}
		avg += (int(v)-avg) / (i+1)
	}
	return int16(avg)
}