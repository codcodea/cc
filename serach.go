package main

import (
	"fmt"
	"strings"

	"github.com/codcodea/cc/palette"
	t "github.com/codcodea/cc/types"

	"github.com/lucasb-eyer/go-colorful"
)

// search.go holds the logic for searching and retrieving colors from the KD trees

// AddNames finds the nearest 5 colors in the name tree and adds them to the response.
// - ref: the reference point (user input)
// - res: i pointer to the response struct to be populated
func AddNames(ref *t.CustomPoint, res *t.Response) {
	nearest := Trees["NAM"].KNN(ref, 5)

	for _, n := range nearest {
		name := n.(t.CustomPoint).Name
		res.Names = append(res.Names, name)
	}
}

// RAL colors
func AddRAL(ref *t.CustomPoint, res *t.Response) {
	nearest := Trees["RAL"].KNN(ref, 1)

	if len(nearest) > 0 {
		res.Conversion.RAL.JSONRecord = extractToJson(nearest[0].(t.CustomPoint))
		res.Conversion.RAL.Distance = calDistance(ref, nearest[0].(t.CustomPoint))
	}
}

// PANTONE colors
func AddPAN(ref *t.CustomPoint, res *t.Response) {
	nearest := Trees["PAN"].KNN(ref, 2)

	fmt.Println("PAN", nearest)

	if len(nearest) > 0 {
		res.Conversion.PAN.JSONRecord = extractToJson(nearest[0].(t.CustomPoint))
		res.Conversion.PAN.Distance = calDistance(ref, nearest[0].(t.CustomPoint))
	}
}

// NCS colors
func AddNCS(ref *t.CustomPoint, res *t.Response) {
	nearest := Trees["NCS"].KNN(ref, 1)

	if len(nearest) > 0 {
		res.Conversion.NCS.JSONRecord = extractToJson(nearest[0].(t.CustomPoint))
		res.Conversion.NCS.Distance = calDistance(ref, nearest[0].(t.CustomPoint))
	}
}

// AddMono creates a monocromatic custom gradient of 5 colors from the reference (user selected) color
// - color: the reference color in HEX format
// - ref: the reference point
// - res: a pointer to the response struct to be populated
func AddMono(color string, ref *t.CustomPoint, res *t.Response) {
	// NatrualGradient from /palette/
	palette, err := palette.NatrualGradient(color)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range palette.Gradient {
		c := &t.Color{
			Color: p.Hex,
			Name:  GetColorName(&p.Lab),
		}
		res.Mono = append(res.Mono, struct{ t.Color }{*c})
	}
}

// *** HELPER FUNCTIONS ***

// calDistance calculates the distance between two colors in LAB color space
// the distance is calculated using the CIEDE2000 algorithm
// CIEDE2000 is industy standard for color difference and lets to end-user evaluate API results
// ref: https://en.wikipedia.org/wiki/Color_difference

func calDistance(ref *t.CustomPoint, nearest t.CustomPoint) float64 {
	c1 := colorful.Lab(ref.Lab.LAB[0], ref.Lab.LAB[1], ref.Lab.LAB[2])
	c2 := colorful.Lab(nearest.Lab.LAB[0], nearest.Lab.LAB[1], nearest.Lab.LAB[2])
	return c1.DistanceCIEDE2000(c2)
}

// helper function to extract name and color from data structures
func extractToJson(nearest t.CustomPoint) t.JSONRecord {
	name := nearest.Name
	lab := nearest.Lab

	record := new(t.JSONRecord)
	record.Name = name
	record.Lab.LABjson.L = lab.LAB[0]
	record.Lab.LABjson.A = lab.LAB[1]
	record.Lab.LABjson.B = lab.LAB[2]

	return *record
}

// GetColorName returns the name of the nearest color
// is uses the same logic as AddNames but returns only the first result
func GetColorName(ref *t.CustomPoint) string {
	nearest := Trees["NAM"].KNN(ref, 1)
	if len(nearest) > 0 {
		return nearest[0].(t.CustomPoint).Name // cast *kdtree.Point to CustomPoint
	}
	return ""
}

// Type definitions for the form session
type ColorfulJson struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type FilterNames struct {
	Names map[string]string `json:"names"`
}

type FormSession struct {
	LastQuery  string
	LastResult []ColorfulJson
}

func NewFormSession() *FormSession {
	return &FormSession{
		LastQuery:  "",
		LastResult: []ColorfulJson{},
	}
}

func ColorLookUp(query string, session *FormSession) {

	if len(query) < 2 {
		return // ignore short queries
	}

	query = strings.ToLower(query);

	isExtended := strings.HasPrefix(query, session.LastQuery)
	isEmpty := session.LastQuery == ""

	words := strings.Split(query, " ")

	// Check if the current query extends the previous query
	if !isExtended || isEmpty {

		session.LastResult = []ColorfulJson{}

		for _, c := range Names {
			match := true
			for _, w := range words {
				if w == "" {
					continue
				}
				if !strings.Contains(c.Name, w) {
					match = false
					break
				}
			}

			if match {
				colorFul := ColorfulJson{c.Name, c.Color.Hex()}
				session.LastResult = append(session.LastResult, colorFul)
			}
		}

		fmt.Println("!isExtended", len(session.LastResult))

	} else {
		// Filter the previous result
		filteredResult := []ColorfulJson{}

		for _, v := range session.LastResult {
			match := true
			for _, w := range words {
				if w == "" {
					continue
				}
				if !strings.Contains(v.Name, w) {
					match = false
					break
				}
			}

			if match {
				filteredResult = append(filteredResult, v)
			}
		}
		session.LastResult = filteredResult
	}
	session.LastQuery = query
}
