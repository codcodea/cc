package io

import (
	"fmt"
	"github.com/kyroy/kdtree"
)


func AddNCS(data *Data, ref kdtree.Point, tree *kdtree.KDTree ) {

	nearest := tree.KNN(ref, 1)

	if len(nearest) > 0 {

		fmt.Println("PANtree:", nearest[0])
		nearestName := nearest[0].(CustomPoint).Name
		fmt.Println("PANtree:", nearestName)
	} else {
		fmt.Println("No nearest neighbor found")
	}
	
}