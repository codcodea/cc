package io

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kyroy/kdtree"
)

var (
	NAM = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/target/colornames.json"
	RAL = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/target/RAL_PLUS_CIELAB1931_sRGB.json"
	PAN = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/target/pantone.json"
	NCS = "/Users/fhalsius/Documents/Code/go/mycolorpicker/io/target/ncs.json"
)


type Point interface {
	Dimensions() int
	Dimension(i int) float64
}

// Custom point type.
type LAB struct {
	LAB [3]float64 `json:"lab"`
}

type JSONRecord struct {
	Name string `json:"name"`
	Lab  struct {
		L float64 `json:"L"`
		A float64 `json:"a"`
		B float64 `json:"b"`
	} `json:"lab"`
}


func NewPoint(name string, p [3]float64) kdtree.Point {
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

func LoadNameKDTree() (*kdtree.KDTree, error){

	start := time.Now()

	// Read the JSON file into a byte slice.
	jsonData, err := os.ReadFile(NAM)
	if err != nil {
		return nil, err
	}

	// Define a slice to hold the parsed JSON records.
	var records []JSONRecord

	// Unmarshal the JSON data into the records slice.
	if err := json.Unmarshal(jsonData, &records); err != nil {
		return nil, err
	}

	// Create a KD tree and populate it with the parsed data.
	tree := kdtree.New(nil)
	for _, record := range records {

		p := [3]float64{record.Lab.L, record.Lab.A, record.Lab.B}
		point := NewPoint(record.Name, p)
		tree.Insert(point)
	}

	end := time.Since(start)

	fmt.Println("LoadNameKDTree:", end)

	return tree, nil
}

func LoadRALKDTree() (*kdtree.KDTree, error){
	
	start := time.Now()

	// Read the JSON file into a byte slice.
	jsonData, err := os.ReadFile(RAL)
	if err != nil {
		return nil, err
	}

	// Define a slice to hold the parsed JSON records.
	var records []JSONRecord

	// Unmarshal the JSON data into the records slice.
	if err := json.Unmarshal(jsonData, &records); err != nil {
		return nil, err
	}

	// Create a KD tree and populate it with the parsed data.
	tree := kdtree.New(nil)
	for _, record := range records {
		p := [3]float64{record.Lab.L, record.Lab.A, record.Lab.B}
		point := NewPoint(record.Name, p)
		tree.Insert(point)
	}

	end := time.Since(start)

	fmt.Println("LoadRALKDTree:", end)

	return tree, nil
}

func LoadPANKDTree() (*kdtree.KDTree, error){
	
	start := time.Now()

	// Read the JSON file into a byte slice.
	jsonData, err := os.ReadFile(PAN)
	if err != nil {
		return nil, err
	}

	// Define a slice to hold the parsed JSON records.
	var records []JSONRecord

	// Unmarshal the JSON data into the records slice.
	if err := json.Unmarshal(jsonData, &records); err != nil {
		return nil, err
	}

	// Create a KD tree and populate it with the parsed data.
	tree := kdtree.New(nil)
	for _, record := range records {
		p := [3]float64{record.Lab.L, record.Lab.A, record.Lab.B}
		point := NewPoint(record.Name, p)
		tree.Insert(point)
	}

	end := time.Since(start)

	fmt.Println("LoadPANKDTreek:", end)

	return tree, nil
}

func LoadNCSKDTree() (*kdtree.KDTree, error){
	
	start := time.Now()

	// Read the JSON file into a byte slice.
	jsonData, err := os.ReadFile(NCS)
	if err != nil {
		return nil, err
	}

	// Define a slice to hold the parsed JSON records.
	var records []JSONRecord

	// Unmarshal the JSON data into the records slice.
	if err := json.Unmarshal(jsonData, &records); err != nil {
		return nil, err
	}

	// Create a KD tree and populate it with the parsed data.
	tree := kdtree.New(nil)
	for _, record := range records {
		p := [3]float64{record.Lab.L, record.Lab.A, record.Lab.B}
		point := NewPoint(record.Name, p)
		tree.Insert(point)
	}

	end := time.Since(start)

	fmt.Println("LoadNCSKDTree:", end)

	return tree, nil
}
