// convert png's to 2d texture arrays for opengl
package main

import (
	"image"
	"log"
	"os"
)
import _ "image/png"
import _ "image/jpeg"

// says http://golang.org/pkg/image/color/#Color
const SIZEOF_UINT32 = 0xFFFF

func img2array(img image.Image) (imgSlice []float32, spanX, spanY int) {
	// Just for documentation:
	// image.Image has it's (0,0) upper left,
	// opengl has its (0,0) lower left

	// opengl array seems to start at 0,1, ... 1,1, 0,0, ..., 1,0
	rect := img.Bounds()
	spanX = rect.Max.X - rect.Min.X
	spanY = rect.Max.Y - rect.Min.Y
	log.Printf("image size: %v x %v", spanX, spanY)
	//imgSlice = make([]float32, spanX*spanY)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			imgSlice = append(
				imgSlice,
				float32(r)/SIZEOF_UINT32,
				float32(g)/SIZEOF_UINT32,
				float32(b)/SIZEOF_UINT32,
				float32(a)/SIZEOF_UINT32,
			)
		}
	}
	log.Print("return from img2array")
	return
}

func png2array(filename string) (imgSlice []float32, width int, height int) {
	imgFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Loading image file failed: %v\n", err)
	}
	img, format, err := image.Decode(imgFile)
	if err != nil {
		log.Fatalf("Decoding image file failed: %v\n", err)
	}
	if format != "png" && format != "jpeg" {
		log.Printf("Strangely, format was not png but %v!\n", format)
	}
	log.Print("png2array retunrs img2array")
	return img2array(img)
}
