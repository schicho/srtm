package main

import (
	"fmt"
	"os"

	"github.com/schicho/srtm"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	var format srtm.SRTMFormat
	fmt.Printf("file size is %d bytes\n", stat.Size())
	switch stat.Size() {
		case 25934402:
			fmt.Println("Assuming SRTM1 format")
			format = srtm.SRTM1Format
		case 2884802:
			fmt.Println("assuming SRTM3 format")
			format = srtm.SRTM3Format
		default:
			fmt.Println("unknown filesize. exiting")
			os.Exit(1)
	}

	srtm, err := srtm.NewSRTMImage(f, format)
	if err != nil {	
		fmt.Println(err)
		os.Exit(1)
	}
	min, max := srtm.MinMaxElevation()
	countDatavoids := srtm.CountDataVoids()
	fmt.Println("min/max values may be erroneous, because of voids or other invalid data")
	fmt.Println("min:", min, "max:", max, "mean:", srtm.MeanElevation())
	fmt.Println("count of data voids:", countDatavoids)
}
