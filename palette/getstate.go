package palette

// GetPointState determines the state of a point within a gradient of 5 colors, ensuring it stays within the specified range.
// For example, if the rangeBounds are [0, 100]:
// - If the reference color is at either end of the range (0 or 100), the function returns -1 or 1, respectively.
//   This indicates that the API will generate a gradient of 5 colors from the reference color towards the opposite end of the range.

// [*] reference color, generated colors[-]

// [*] [-] [-] [-] [-]
// [-] [*] [-] [-] [-]
// [-] [-] [*] [-] [-]
// [-] [-] [-] [*] [-]
// [-] [-] [-] [-] [*]

//GetPointState is a port of the original JavaScript code "getstate.js."
func GetPointState(p3 float64, pointSpacing float64, rangeBounds [2]float64) int {
	numPoints := 5
	span := pointSpacing * float64(numPoints-1)

	p1 := p3 - span/2
	p2 := p3 - span/4
	p4 := p3 + span/4
	p5 := p3 + span/2

	if p1 < rangeBounds[0] && p2 < rangeBounds[0] {
		return -2
	} else if p1 < rangeBounds[0] {
		return -1
	} else if p5 > rangeBounds[1] && p4 > rangeBounds[1] {
		return 2
	} else if p5 > rangeBounds[1] {
		return 1
	} else {
		return 0
	}
}
