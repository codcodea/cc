package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/codcodea/cc/types"
	"github.com/kyroy/kdtree"
	"github.com/lucasb-eyer/go-colorful"
)

type FileName struct {
	Name string
	Path string
}

type Colorful struct {
	Name  string
	Color colorful.Color
}

var (
	NAM = FileName{"NAM", "/Users/fhalsius/Documents/Code/go/mycolorpicker/db/target/colornames.json"}
	RAL = FileName{"RAL", "/Users/fhalsius/Documents/Code/go/mycolorpicker/db/target/RAL_PLUS_CIELAB1931_sRGB.json"}
	PAN = FileName{"PAN", "/Users/fhalsius/Documents/Code/go/mycolorpicker/db/target/pantone.json"}
	NCS = FileName{"NCS", "/Users/fhalsius/Documents/Code/go/mycolorpicker/db/target/ncs.json"}
)

var (
	Trees = make(map[string]*kdtree.KDTree) // map of all serach trees
	Names = []Colorful{}                    // map of all color names <name, hex>
)

// init.go holds the initialization logic for the server

// When the server starts:
// - each tree is loaded from a JSON file
// - each tree is stored as a binary KD tree 
// - each tree is stored in memory
// - on subsequent requests the memory trees are used for search and retrieval

// KD tree implementation:
// General info Wiki: https://en.wikipedia.org/wiki/K-d_tree
// Library and methods: https://github.com/kyroy/kdtree

func LoadTrees() error {
	filePaths := []FileName{NAM, RAL, PAN, NCS}

	for _, f := range filePaths {
		tree, err := LoadColorTree(f.Path)
		if err != nil {
			fmt.Println("Error rebuilding KD tree:", err)
			return err
		}
		Trees[f.Name] = tree
	}
	return nil
}

func LoadColorTree(filePath string) (*kdtree.KDTree, error) {
	// Read and parse JSON
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var records []types.JSONRecord

	if err := json.Unmarshal(jsonData, &records); err != nil {
		return nil, err
	}
	// Create and populate a KD tree
	tree := kdtree.New(nil)
	for _, record := range records {
		p := [3]float64{record.Lab.L, record.Lab.A, record.Lab.B}
		point := types.NewPoint(record.Name, p)
		tree.Insert(point)
	}
	return tree, nil
}

func LoadNameMap() error {
	var fileNAME = "/Users/fhalsius/Documents/Code/go/mycolorpicker/db/source/colornames.csv"

	csvFile, err := os.Open(fileNAME)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	_, err = reader.Read() // skip first line
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return err
	}

	for {
		line, err := reader.Read()
		if err != nil {
			break // End of file.
		}

		// Parse
		nam := strings.ToLower(line[0])
		hex := strings.ToLower(line[1])

		// Load color into a colorful.Color struct
		c, err := colorful.Hex(string(hex))
		if err != nil {
			fmt.Println("Load Names conversion err:", err)
			break
		}

		// Store in map
		Names = append(Names, Colorful{nam, c})
	}

	// Pre sort the names
	sort.Slice(Names, func(i, j int) bool {
		l1, c1, h1 := Names[i].Color.LuvLCh()
		l2, c2, h2 := Names[j].Color.LuvLCh()
		if l1 != l2 {
			return l1 < l2
		}
		if h1 != h2 {
			return h1 < h2
		}
		if c1 != c2 {
			return c1 < c2
		}
		return false
	})

	slices.Reverse(Names);
	fmt.Println("Color names loaded:", len(Names))
	return nil
}
