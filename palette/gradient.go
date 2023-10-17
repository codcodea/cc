package palette

import (
	"errors"
	"math"

	"github.com/codcodea/cc/types"
	"github.com/lucasb-eyer/go-colorful"
)

// ** Type definitions

type Gradient struct {
	Hex string
	Lab types.CustomPoint
}

type Gradients struct {
	Gradient []Gradient
}

// NaturalGradient generates a natural-looking gradient based on the input color.
// Parameters:
// - hex (string): The hexadecimal color code.
// Returns:
// - Gradients: A struct containing 5 colors representing the gradient.

// In nature, gradients exhibit non-linear changes in hue, saturation, and lightness.
// Digital gradients, in contrast, are typically linear and mainly alter the lightness.

// This function incorporates "magic numbers" derived from the author's measurements, experience,
// and heuristics. It can be regarded as the "secret sauce" for creating natural-looking gradients.


func NatrualGradient(hex string) (Gradients, error) {

	// toHSL
	c, err := colorful.Hex(hex)

	if err != nil {
		return Gradients{}, errors.New("error converting hex to HSL")
	}

	h, s, l := c.Hsl()

	// Get the state of the l ligntness point in the a gradient [0-1]
	state := GetPointState(l, 0.2, [2]float64{0.0, 1.0})

	var hue []float64
	var sat []float64
	var lit []float64

	litHighBreakPoint := 0.79
	//litLowBreakPoint := 0.5 // not used

	satLowBreakPoint := 0.08
	satHighBreakPoint := 0.5

	lightBooster := 0.0
	saturationBooster := 0.0

	if s < satLowBreakPoint {
		lightBooster = 0.03
		saturationBooster = 0.05
	} else if s > satHighBreakPoint {
		lightBooster = 0.0
		saturationBooster = -0.05
	}

	if l < litHighBreakPoint {
		lightBooster = 0.0
	} else if l > litHighBreakPoint {
		lightBooster = -0.03
	}

	satCorr := saturationBooster
	lightCorr := lightBooster

	litV := 0.06

	if state == 0 {
		litV = 0.09
	}

	if state == 1 || state == -1 {
		litV = 0.07
	}

	if state == 0 {
		hue = []float64{h - 2, h - 1, h, h + 1, h + 2}
		sat = []float64{s * 0.83, s * 0.91, s, s * (1.1 + satCorr), s * (1.2 + satCorr*2)}
		lit = []float64{(l - litV*1.75) + 2*lightCorr, (l - litV) + lightCorr, l, (l + litV) + lightCorr, (l + 2*litV) + 2*lightCorr}
	} else if state == 1 {
		hue = []float64{h - 3, h - 2, h - 1, h, h + 1}
		sat = []float64{s * 0.77, s * 0.83, s * 0.91, s, s * (1.1 + satCorr)}
		lit = []float64{(l - 3*litV) + 3*lightCorr, (l - 2*litV) + 2*lightCorr, (l - litV) + lightCorr, l, (l + litV) + lightCorr}
	} else if state == 2 {
		hue = []float64{h - 4, h - 3, h - 2, h - 1, h}
		sat = []float64{s * 0.5, s * 0.6, s * 0.7, s * 0.8, s}
		lit = []float64{l - 4*litV, l - 3*litV, l - 2*litV, l - litV, l}
	} else if state == -2 {
		hue = []float64{h, h + 1, h + 2, h + 3, h + 4}
		sat = []float64{s, s, s, s * (1.1 + satCorr), s * (1.2 + satCorr*2)}
		lit = []float64{l, l + litV + lightCorr, l + 2*litV + 2*lightCorr, (l + 3*litV) + 3*lightCorr, (l + 4*litV) + 4*lightCorr}
	} else if state == -1 {
		hue = []float64{h - 1, h, h + 1, h + 2, h + 3}
		sat = []float64{s * 0.9, s, s * (1.1 + satCorr), s * (1.2 + satCorr*2), s * (1.3 + satCorr*2)}
		lit = []float64{l - litV, l, l + litV, (l + litV*2) + 2*lightCorr, (l + litV*3) + 3*lightCorr}
	}

	// Construct the return struct
	gradients := new(Gradients)

	for i := 0; i < 5; i++ {
		var adjHue, adjLit, adjSat float64

		adjHue = math.Mod(hue[i], 360.0)
		if adjHue < 0 {
			adjHue += 360.0
		}

		adjLit = math.Max(0.01, math.Min(0.99, lit[i]))
		adjSat = math.Max(0.0, math.Min(1.0, sat[i]))

		// toHexLab
		c := colorful.Hsl(adjHue, adjSat, adjLit)
		hex := c.Hex()
		l, a, b := c.Lab()

		// Bind
		gradient := new(Gradient)
		gradient.Hex = hex
		gradient.Lab = types.NewPoint("ref", [3]float64{l, a, b})

		gradients.Gradient = append(gradients.Gradient, *gradient)
	}

	return *gradients, nil
}
