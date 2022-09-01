package srtm

import (
	"encoding/binary"
	"io"
)

var SRTMByteOrder = binary.BigEndian

type SRTMFormat int

const (
	SRTM1Format = SRTMFormat(iota)
	SRTM3Format
)

const (
	SRTM1Size = 3601
	SRTM3Size = 1201
)

func (f SRTMFormat) Size() int {
	switch f {
	case SRTM1Format:
		return SRTM1Size
	case SRTM3Format:
		return SRTM3Size
	}
	return -1
}

type SRTMImage struct {
	Data []int16
	Format SRTMFormat
}


func NewSRTMImage(r io.Reader, format SRTMFormat) (*SRTMImage, error) {
	data := make([]int16, format.Size() * format.Size())
	err := binary.Read(r, SRTMByteOrder, data)
	return &SRTMImage{data, format}, err
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
// Overflows are mitigated. Thus may not be the actual mean.
func (s *SRTMImage) MeanElevation() int16 {
	var avg = 0

	for i, v := range s.Data {
		avg += (int(v)-avg) / (i+1)
	}
	return int16(avg)
}