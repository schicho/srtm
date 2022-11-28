package main

import (
	"log"
	"os"

	"github.com/schicho/srtm"
	"golang.org/x/image/tiff"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var format srtm.SRTMFormat
	stat, err := f.Stat()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	switch stat.Size() {
	case 25934402:
		log.Println("Assuming SRTM1 format")
		format = srtm.SRTM1Format
	case 2884802:
		log.Println("assuming SRTM3 format")
		format = srtm.SRTM3Format
	default:
		log.Println("unknown filesize. exiting")
		os.Exit(1)
	}

	srtmImg, err := srtm.NewSRTMImage(f, format)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	f_out, err := os.Create(stat.Name() + "-out-16.tiff")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f_out.Close()

	err = tiff.Encode(f_out, srtmImg.FullImage(), nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("Please note that the bit-depth is 16 bit per pixel.\nSome viewers may not be able to display the image.")
}
