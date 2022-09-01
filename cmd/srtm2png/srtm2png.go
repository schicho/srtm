package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/schicho/srtm"
)

func main() {
	file := flag.String("i", "", "specify the SRTM file")
	resolution := flag.Int("r", 0, "specify the resolution of the SRTM data")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	if !(*resolution == 1 || *resolution == 3) {
		flag.Usage()
		fmt.Println("resolution must be 1 or 3")
		os.Exit(1)
	}

	f_in, err := os.Open(*file) 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f_in.Close()

	var format srtm.SRTMFormat
	switch *resolution {
	case 1:	format = srtm.SRTM1Format
	case 3:	format = srtm.SRTM3Format
	}

	srtm, err := srtm.NewSRTMImage(f_in, format)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	f_out, err := os.Create("out.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f_out.Close()
	
	err = png.Encode(f_out, srtm.MeanCenteredImage())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}