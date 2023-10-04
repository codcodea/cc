package types

type Response struct {
	Base struct {
		Color string `json:"color"`
		Name  string `json:"name"`
	} `json:"base"`

	Mono []struct {
		Color string `json:"color"`
		Name  string `json:"name"`
	} `json:"mono"`

	Names []string `json:"names"`

	Conversion struct {
		RGB  string `json:"rgb"`
		HSL  string `json:"hsl"`
		HSV  string `json:"hsv"`
		LAB  string `json:"lab"`
		CMYK string `json:"cmyk"`
		RAL  struct {
			Name string `json:"name"`
			LABjson
			Distance float64 `json:"distance"`
		} `json:"ral"`
		Pantone struct {
			Name string `json:"name"`
			LABjson
			Distance float64 `json:"distance"`
		} `json:"pantone"`
		NCS struct {
			Name string `json:"name"`
			LABjson
			Distance float64 `json:"distance"`
		} `json:"ncs"`
	} `json:"conversions"`
}

type LAB struct {
	Lab [3]float64
}

type LABjson struct {
	L float64 `json:"L"`
	A float64 `json:"a"`
	B float64 `json:"b"`
}
