package io

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

var nameSourcePath = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/source/colornames.csv"
var nameTargetPath = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/target/colornames.json"

var ralSourcePath = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/source/RAL_PLUS_CIELAB1931_sRGB.csv"
var ralTargetPath = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/target/RAL_PLUS_CIELAB1931_sRGB.json"

var palSourcePath = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/source/pantone.txt"
var palTargetPath = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/target/pantone.json"

var nscSourcePath = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/source/ncs.txt"
var nscTargetPath = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/target/ncs.json"

// Target structure for JSON data.
type ColorData struct {
	Name string `json:"name"`
	Lab  struct {
		L float64 `json:"L"`
		A float64 `json:"a"`
		B float64 `json:"b"`
	} `json:"lab"`
	Code *string `json:"code,omitempty"`
}

// Source structure for RAL CSV data.
type RALColorData struct {
	Name      string  `json:"name"`
	Hue       string  `json:"hue"`
	Lightness string  `json:"lightness"`
	Chroma    string  `json:"chroma"`
	RGB       [3]int  `json:"rgb"`
	Code      *string `json:"code,omitempty"`
}

type NCSColorData struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

func Clean() {

	csvFile, err := os.Open(nameSourcePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	var colorDataSlice []ColorData

	for {

		line, err := reader.Read()
		if err != nil {
			break // End of file.
		}

		// Parse
		name := line[0]
		hexValue := line[1]

		// Hex to Lab
		color, _ := colorful.Hex(hexValue)
		l, a, b := color.Lab()

		// Store
		colorData := ColorData{
			Name: name,
			Lab: struct {
				L float64 `json:"L"`
				A float64 `json:"a"`
				B float64 `json:"b"`
			}{
				L: l,
				A: a,
				B: b,
			},
		}

		// Append the ColorData to the slice.
		colorDataSlice = append(colorDataSlice, colorData)
	}

	// Convert color data slice to JSON.
	jsonData, err := json.Marshal(colorDataSlice)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON data to a JSON file.
	jsonFile, err := os.Create(nameTargetPath)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Println("Conversion completed. JSON file created.")
}

func CleanRAL() {
	// Open the CSV file for reading.
	csvFile, err := os.Open(ralSourcePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV reader.
	reader := csv.NewReader(csvFile)
	reader.Comma = ';'

	// Initialize a slice to store RAL color data.
	var ralData []RALColorData

	for {
		// Read a line from the CSV.
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// End of file reached.
				break
			}
			fmt.Println("Error reading CSV row:", err)
			return
		}

		// Extract RAL color data from the CSV.
		hue := row[1]
		lightness := row[2]
		chroma := row[3]
		r, _ := strconv.Atoi(row[4])
		g, _ := strconv.Atoi(row[5])
		b, _ := strconv.Atoi(row[6])
		code := row[7]

		// Create a RALColorData structure.
		colorData := RALColorData{
			Name:      code,
			Hue:       hue,
			Lightness: lightness,
			Chroma:    chroma,
			RGB:       [3]int{r, g, b},
		}

		// Append the RALColorData to the slice.
		ralData = append(ralData, colorData)
	}

	// Convert RGB to LAB.
	labData := ralRGBtoLab(ralData)

	// Convert labData to JSON.
	jsonData, err := json.Marshal(labData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON data to a JSON file.
	jsonFile, err := os.Create(ralTargetPath)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Println("Conversion completed. JSON file created: ral_colors.json")
}

func ralRGBtoLab(ralData []RALColorData) []ColorData {
	var labData []ColorData

	for _, entry := range ralData {
		rgb := colorful.Color{
			R: float64(entry.RGB[0]) / 255.0,
			G: float64(entry.RGB[1]) / 255.0,
			B: float64(entry.RGB[2]) / 255.0,
		}

		l, a, b := rgb.Lab()

		colorData := ColorData{
			Name: entry.Name,
			Lab: struct {
				L float64 `json:"L"`
				A float64 `json:"a"`
				B float64 `json:"b"`
			}{
				L: l,
				A: a,
				B: b,
			},
		}

		labData = append(labData, colorData)
	}

	return labData
}

func CleanPAN() {
	// Open the source text file for reading.
	file, err := os.Open(palSourcePath)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer file.Close()

	// Create a slice to store the color data.
	var colorDataSlice []ColorData

	// Create a scanner to read lines from the file.
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file.
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by space to extract RGB values and Pantone code.
		parts := strings.Fields(line)
		if len(parts) != 4 {
			fmt.Println("Error parsing line:", line)
			continue
		}

		// Extract RGB values and Pantone code.
		r, _ := strconv.Atoi(parts[0])
		g, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		pantoneCode := parts[3]

		// Calculate LAB values from RGB.

		c := colorful.Color{float64(r), float64(g), float64(b)}
		rr, gg, bb := c.Lab()

		// Create a ColorData struct and add it to the slice.
		colorData := ColorData{
			Name: pantoneCode,
			Lab: struct {
				L float64 `json:"L"`
				A float64 `json:"a"`
				B float64 `json:"b"`
			}{
				L: rr,
				A: gg,
				B: bb,
			},
		}
		colorDataSlice = append(colorDataSlice, colorData)
	}

	// Check for scanner errors.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading source file:", err)
		return
	}

	// Serialize the color data to JSON.
	colorDataJSON, err := json.Marshal(colorDataSlice)
	if err != nil {
		fmt.Println("Error marshaling color data to JSON:", err)
		return
	}

	// Write the JSON data to a file.
	if err := os.WriteFile(palTargetPath, colorDataJSON, os.ModePerm); err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Println("Color data converted and saved")
}

func CleanNCS() {
	// Open the source text file for reading.
	file, err := os.Open(nscSourcePath)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer file.Close()

	// Create a slice to store the color data.
	var colorDataSlice []ColorData

	// Create a scanner to read lines from the file.
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file.
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimPrefix(line, "$")
		line = strings.TrimSpace(line)

		// Split the line by space to extract RGB values and Pantone code.
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Error parsing line:", line)
			continue
		}

		// Extract RGB values and Pantone code.
		name := parts[0]
		hex := parts[1]

		// Calculate LAB values from RGB.

		c, err := colorful.Hex(hex)

		if err != nil {
			fmt.Println("Error parsing hex value:", hex)
			break
		}

		rr, gg, bb := c.Lab()

		// Create a ColorData struct and add it to the slice.
		colorData := ColorData{
			Name: name,
			Lab: struct {
				L float64 `json:"L"`
				A float64 `json:"a"`
				B float64 `json:"b"`
			}{
				L: rr,
				A: gg,
				B: bb,
			},
		}
		colorDataSlice = append(colorDataSlice, colorData)
	}

	// Check for scanner errors.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading source file:", err)
		return
	}

	// Serialize the color data to JSON.
	colorDataJSON, err := json.Marshal(colorDataSlice)
	if err != nil {
		fmt.Println("Error marshaling color data to JSON:", err)
		return
	}

	// Write the JSON data to a file.
	if err := os.WriteFile(nscTargetPath, colorDataJSON, os.ModePerm); err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Println("Color data converted and saved")
}
