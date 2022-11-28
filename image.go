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
	mean := srtmImg.ElevationMean()
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
// This preserves the original data completely and allows for further processing.
//
// Please note that some image viewers are not able to display 16 bit images.
// In many cases, the image will be displayed as a 8 bit image or will be displayed incorrectly.
func (srtmImg *SRTMImage) FullImage() *image.Gray16 {
	rect := image.Rect(0, 0, srtmImg.Format.Size(), srtmImg.Format.Size())
	img := image.NewGray16(rect)

	for i, v := range srtmImg.Data {
		// Gray16 uses uint16, but the data is int16, so we shift the values by 32768
		// Gray16 uses big endian, according to its documentation
		binary.BigEndian.PutUint16(img.Pix[i*2:i*2+2], signed16BitToUint16(v))
	}
	return img
}

// signed16BitToUint16 converts a signed 16 bit integer to an unsigned 16 bit integer.
// It shifts the value by 32768. e.g. -32768 becomes 0, 0 becomes 32768, 32767 becomes 65535.
func signed16BitToUint16(v int16) uint16 {
	return uint16((int32(v) + 32768) & 0xFFFF)
}
