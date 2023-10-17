package io

import (
	"fmt"

	ty "github.com/codcodea/cc/types"
	"github.com/lucasb-eyer/go-colorful"
)

type Conversion struct {
	RGB  string
	HSL  string
	HSV  string
	LAB  string
	CMYK string
}

func Convert(hex string, res *ty.Response) (ty.CustomPoint, error) {

	conv := &res.Conversion
	res.Base.Color.Color = hex

	c, err := colorful.Hex(hex)

	if err != nil {
		return ty.CustomPoint{}, err
	}

	r, g, b := c.RGB255()
	h, s, l := c.Hsl()
	hh, ss, v := c.Hsv()
	ll, aa, bb := c.Lab()
	str := rgbToCmyk(r, g, b)

	conv.RGB = fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
	conv.HSL = fmt.Sprintf("hsl(%.1f, %.1f%%, %.1f%%)", h, s*100, l*100)
	conv.HSV = fmt.Sprintf("hsv(%.1f, %.1f%%, %.1f%%)", hh, ss*100, v*100)
	conv.LAB = fmt.Sprintf("lab(%.2f, %.2f, %.2f)", ll, aa, bb)
	conv.CMYK = str

	ref :=  ty.NewPoint("ref", [3]float64{ll, aa, bb})

	return ref, nil

}

func rgbToCmyk(r, g, b uint8) string {
	// Convert RGB values to the range [0, 1]
	rFloat := float64(r) / 255.0
	gFloat := float64(g) / 255.0
	bFloat := float64(b) / 255.0

	// Calculate the CMY components
	c := 1 - rFloat
	m := 1 - gFloat
	y := 1 - bFloat

	// Find the minimum of CMY values
	minCMY := min(min(c, m), y)

	// Avoid division by zero and calculate K (black)
	k := 1.0
	if minCMY < 1.0 {
		k = minCMY
		c = (c - k) / (1 - k)
		m = (m - k) / (1 - k)
		y = (y - k) / (1 - k)
	}

	return fmt.Sprintf("cmyk(%.1f%%, %.1f%%, %.1f%%, %.1f%%)", c*100, m*100, y*100, k*100)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}