package main

import (
	"image/png"
	"log"
	"os"

	"github.com/schicho/srtm"
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

	f_out, err := os.Create(stat.Name() + "-out-8.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f_out.Close()

	p4 := srtmImg.ElevationPercentile(0.04)
	p95 := srtmImg.ElevationPercentile(0.95)
	err = png.Encode(f_out, srtmImg.ScaledHeightImage(2, int16((p4+p95)/2)))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
