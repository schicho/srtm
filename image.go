package srtm

import (
	"encoding/binary"
	"image"
)

// MeanCenteredImage centeres the elevation data to the mean elevation value.
// No scaling is applied.
// Values may be erroneous, because of voids or other invalid data.
func (s *SRTMImage) MeanCenteredImage() *image.Gray {
	mean := s.MeanElevation()
	return centerHeightImage(s.Data, s.Format, mean)
}

// CenterHeightImage centeres the elevation data to the provided elevation value.
// Values may be erroneous, because of voids or other invalid data.
// No scaling is applied.
// Takes a center value as input. This value will be centered to 128.
// Large or small values may result in overflows or underflows.
func (s *SRTMImage) CenterHeightImage(center int16) *image.Gray {
	return centerHeightImage(s.Data, s.Format, center)
}

// centerHeightImage centeres the elevation data to the provided elevation value.
func centerHeightImage(data []int16, format SRTMFormat, center int16) *image.Gray {
	rect := image.Rect(0, 0, format.Size(), format.Size())
	img := image.NewGray(rect)

	for i, v := range data {
		// center the data to 128.
		centeredValue := v - center + 128
		if centeredValue < 0 {
			centeredValue = 0
		} else if centeredValue > 255 {
			centeredValue = 255
		}
		img.Pix[i] = uint8(centeredValue)
	}
	return img
}

// FullImage returns the full data range as a 16 bit-depth grayscale image.
// Please note that some image viewers may not display the full range of values.
// Or even display wrong values.
func (s *SRTMImage) FullImage() *image.Gray16 {
	rect := image.Rect(0, 0, s.Format.Size(), s.Format.Size())
	img := image.NewGray16(rect)

	for i, v := range s.Data {
		var converted uint16 = uint16((int32(v) + 32768) & 0xFFFF)
		binary.BigEndian.PutUint16(img.Pix[i*2:i*2+2], converted)
	}
	return img
}
