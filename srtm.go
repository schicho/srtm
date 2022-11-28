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
	Data   []int16
	Format SRTMFormat
}

func NewSRTMImage(r io.Reader, format SRTMFormat) (*SRTMImage, error) {
	data := make([]int16, format.Size()*format.Size())
	err := binary.Read(r, SRTMByteOrder, data)
	return &SRTMImage{data, format}, err
}
