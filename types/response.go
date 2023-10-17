package types

// Type definition for the API response

type Color struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

type Response struct {
	Base struct {
		Color 
	} `json:"base"`

	Mono []struct {
		Color 
	} `json:"mono"`

	Names []string `json:"names"`

	Conversion struct {
		RGB  string `json:"rgb"`
		HSL  string `json:"hsl"`
		HSV  string `json:"hsv"`
		LAB  string `json:"lab"`
		CMYK string `json:"cmyk"`
		RAL  struct {
			JSONRecord
			Distance float64 `json:"distance"`
		} `json:"ral"`
		PAN struct {
			JSONRecord
			Distance float64 `json:"distance"`
		} `json:"pan"`
		NCS struct {
			JSONRecord
			Distance float64 `json:"distance"`
		} `json:"ncs"`
	} `json:"conversions"`
}

