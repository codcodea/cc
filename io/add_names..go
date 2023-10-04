package io

import (
	"fmt"

	"github.com/kyroy/kdtree"
)


func AddNames(data *Data, ref kdtree.Point, tree *kdtree.KDTree ) string {

	nearest := tree.KNN(ref, 1)

	if len(nearest) > 0 {
		nearestName := nearest[0].(CustomPoint).Name
		return nearestName
	} else {
		fmt.Println("No nearest neighbor found")
		return ""
	}
	
}
// nearest := tree.KNN(ref, 5)

// 	for _, n := range nearest {
// 		d.Names = append(d.Names, n.(CustomPoint).Name)
// 	}

func AddName(ref kdtree.Point, tree *kdtree.KDTree ) string {

	nearest := tree.KNN(ref, 1)

	if len(nearest) > 0 {
		nearestName := nearest[0].(CustomPoint).Name
		return nearestName
	} else {
		fmt.Println("No nearest neighbor found")
		return ""
	}
	
}