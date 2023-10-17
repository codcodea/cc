
package types

// Type definition for color and KD-tree

type Point interface {
	Dimensions() int
	Dimension(i int) float64
}

// Custom point type.
type LAB struct {
	LAB [3]float64 `json:"lab"`
}

type LABjson struct {
	L float64 `json:"L"`
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type JSONRecord struct {
	Name string `json:"name"`
	Lab  struct {
		LABjson
	} `json:"lab"`
}


func NewPoint(name string, p [3]float64) CustomPoint {
	return CustomPoint{
		Name: name,
		Lab:  LAB{LAB: p},
	}
}

type CustomPoint struct {
	Name string
	Lab  LAB
}

func (p CustomPoint) Dimensions() int {
	return 3
}

func (p CustomPoint) Dimension(i int) float64 {
	switch i {
	case 0:
		return p.Lab.LAB[0]
	case 1:
		return p.Lab.LAB[1]
	case 2:
		return p.Lab.LAB[2]
	default:
		panic("invalid dimension")
	}
}