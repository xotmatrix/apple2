package apple2

import (
	"image"
	"image/color"
)

const (
	loResPixelWidth       = charWidth * 2
	doubleLoResPixelWidth = charWidth
	loResPixelHeight      = charHeight / 2
)

func getColorPatterns(light color.Color) [16][16]color.Color {
	/*
		For each lores pixel we have to fill 14 half mono pixels with
		the 4 bits of the color repeated. We will need to shift by 2 bits
		on the odd columns. Lets prepare 14+2 values for each color.
	*/

	var data [16][16]color.Color

	for ci := 0; ci < 16; ci++ {
		for cb := uint8(0); cb < 4; cb++ {
			bit := (ci >> cb) & 1
			var colour color.Color
			if bit == 0 {
				colour = color.Black
			} else {
				colour = light
			}
			for i := uint8(0); i < 4; i++ {
				data[ci][cb+4*i] = colour
			}
		}
	}
	return data

}

func snapshotLoResModeMono(a *Apple2, isDoubleResMode bool, isSecondPage bool, isMixMode bool, light color.Color) *image.RGBA {
	text, columns, lines := getActiveText(a, isDoubleResMode, isSecondPage, false)
	if isMixMode {
		lines -= textLinesMix
	}
	grLines := lines * 2
	pixelWidth := loResPixelWidth
	if isDoubleResMode {
		pixelWidth = doubleLoResPixelWidth
	}

	size := image.Rect(0, 0, columns*pixelWidth, grLines*loResPixelHeight)
	img := image.NewRGBA(size)

	patterns := getColorPatterns(light)
	for l := 0; l < grLines; l++ {
		for c := 0; c < columns; c++ {
			char := text[(l/2)*columns+c]
			grPixel := char >> 4
			if l%2 == 0 {
				grPixel = char & 0xf
			}
			// We place pixelWidth mono pixels per graphic pixel.
			// The groups of 4 mono pixels need to be alligned with an offset to get plain surfaces
			offset := (c * pixelWidth) % 4

			// Insert the pixelWidth pixels required
			for i := 0; i < pixelWidth; i++ {
				v := patterns[grPixel][i+offset]
				// Repeat the same color for 4 lines
				for r := 0; r < loResPixelHeight; r++ {
					img.Set(c*loResPixelWidth+i, l*4+r, v)
				}
			}

		}
	}

	return img
}
