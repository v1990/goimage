package goimage

import (
	"github.com/nfnt/resize"
	"image"
	//
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func Resize(img image.Image, width uint, height uint) image.Image {
	m := resize.Resize(width, height, img, resize.Lanczos3)
	return m
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
