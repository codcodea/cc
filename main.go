package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// myColorPicker is the backend engine for myColorPicker.com
// This tool is designed for use by designers and industry professionals working with color.

// The tool provides insights for common challenges when working with color on the web in an easily accessible way:
// - Finding accurate color names for briefs and presentations
// - Expressing colors in industry-standard color systems (RAL, Pantone, and NCS) commonly used in fashion, interior design, and industrial design
// - Expressing colors in CSS standard formats (HEX, RGB, HSL)
// - Serving as a creative tool for color exploration and inspiration

// Use case:
// - The end user browses the internet to find color inspiration for their next project.
// - The user selects a color using the Eye Dropper API from anywhere on the screen.
// - The selected color is sent to the backend.
// - The backend returns the color name, system-specific color codes, and a gradient of 5 colors within the same color family.

// Challenges, motivations, and solutions:
// - Each color is interpolated in LAB among 50,000 color records expressed in the LAB color space (resulting in 300,000 rows of data).
// - Then, each color is reiterated 5 times, resulting in a total of 1,500,000 iterations.
// - By utilizing Golang and KD trees, the search and retrieval are performed in ~1ms, in contrast to the current Node Express implementation which takes 200ms+.

// Main is the entry point of the application.
func main() {

	// Load database of colors into a KD tree into memory on startup
	if err := LoadTrees(); err != nil {
		fmt.Println("Error loading KD trees:", err)
		return
	}

	if err := LoadNameMap(); err != nil {
		fmt.Println("Error loading color names:", err)
		return
	}

	// REST API
	app := echo.New()

	// app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // add mycolorpicker.com in production 
	}))

	app.GET("/", HandleRoot)
	app.GET("/colors/:hex", HandleColor)
	app.GET("/form", HandleLookup)
	app.Logger.Fatal(app.Start(":4005"))

	// Maintainace scripts for the color databases commented out 
	// pk.CleanNAME();
	// return 
}
