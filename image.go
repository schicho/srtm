package srtm

import (
	"encoding/binary"
	"image"
)

// MeanCenteredImage centeres the elevation data to the mean elevation value.
// No scaling is applied. The calculated mean value corresponds to the value of 128 in the
// resulting image. Smaller values are darker, larger values are brighter.
// If there are values too large or too small, these values will be set to white or black, respectively.
// Values may be erroneous, because of voids or other invalid data.
func (srtmImg *SRTMImage) MeanCenteredImage() *image.Gray {
	mean := srtmImg.MeanElevation()
	return heightCenteredImage(srtmImg, mean)
}

// HeightCenteredImage centeres the elevation data to the provided elevation value.
// No scaling is applied. The center value corresponds to the value of 128 in the
// resulting image. Smaller values are darker, larger values are brighter.
// If there are values too large or too small, these values will be set to white or black, respectively.
// Values may be erroneous, because of voids or other invalid data.
func (srtmImg *SRTMImage) HeightCenteredImage(height int16) *image.Gray {
	return heightCenteredImage(srtmImg, height)
}

// heightCenteredImage centeres the elevation data to the provided elevation value.
func heightCenteredImage(srtmImg *SRTMImage, height int16) *image.Gray {
	rect := image.Rect(0, 0, srtmImg.Format.Size(), srtmImg.Format.Size())
	img := image.NewGray(rect)

	height32 := int32(height)

	for i, v := range srtmImg.Data {
		// avoid overflows by using int32
		// center value in the output image is 128
		centeredValue := int32(v) - height32 + 128
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
