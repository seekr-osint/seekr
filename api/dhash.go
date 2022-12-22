package api

import (
	"image"
	"image/color"
	"log"
)

// DHash calculates the dHash of an image
func DHash(img image.Image) (hash uint64) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var lastColor uint8
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			currColor := uint8(r >> 8)
			if x > 0 {
				if currColor > lastColor {
					hash |= 1
				}
				hash <<= 1
			}
			lastColor = currColor
		}
	}
	return
}

func NormalizeImage(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			srcX := (x * bounds.Max.X) / width
			srcY := (y * bounds.Max.Y) / height
			newImg.Set(x, y, img.At(srcX, srcY))
		}
	}
	return newImg
}

// ConvertToGrayscale converts an image to grayscale
func ConvertToGrayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	grayImg := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24)
			grayImg.Set(x, y, color.Gray{gray})
		}
	}
	return grayImg
}
func MkImgHash(img image.Image) uint64 {
	// Normalize the image to a fixed width and height of 4 pixels
	normImg := NormalizeImage(img, 4, 4)

	// Convert the image to grayscale
	grayImg := ConvertToGrayscale(normImg)

	// Calculate the dHash of the grayscale image
	hash := DHash(grayImg)
	log.Printf("Img hash: %d", hash)
	return hash
}
