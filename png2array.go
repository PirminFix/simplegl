// convert png's to 2d texture arrays for opengl
package main

import (
	"image"
	"log"
	"os"
)
import _ "image/png"

func img2array(img image.Image) (imgSlice []uint32, spanX, spanY int) {
	rect := img.Bounds()
	spanX = rect.Max.X - rect.Min.X
	spanY = rect.Max.Y - rect.Min.Y
	imgSlice = make([]uint32, spanX*spanY)
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			imgSlice = append(imgSlice, r, g, b, a)
		}
	}
	return
}

func png2array(filename string) ([]uint32, int, int) {
	imgFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Loading image file failed: %v\n", err)
	}
	img, format, err := image.Decode(imgFile)
	if err != nil {
		log.Fatalf("Decoding image file failed: %v\n", err)
	}
	if format != "png" {
		log.Printf("Strangely, format was not png but %v!\n", format)
	}
	return img2array(img)
}
