package main

import (
	"fmt"

	"github.com/kyroy/kdtree"

	pk "github.com/codcodea/cc/io"
	ty "github.com/codcodea/cc/types"
)

var (
	NameTree *kdtree.KDTree
	RALtree  *kdtree.KDTree
	PANtree  *kdtree.KDTree
	NCStree  *kdtree.KDTree
)

func main() {

	var res ty.Response

	res.Base.Color = "#8A4578"
	res.Base.Name = "Sister Loganberry Frost"

	ref, err := pk.Convert("#edf246", &res)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	NameTree, err = pk.LoadNameKDTree()
	if err != nil {
		fmt.Println("Error rebuilding KD tree from binary:", err)
		return
	}
	pk.AddNames(res, ref, NameTree)

	RALtree, err = pk.LoadRALKDTree()

	if err != nil {
		fmt.Println("Error rebuilding KD tree from binary:", err)
		return
	}
	pk.AddRal(res, ref, RALtree)

	PANtree, err = pk.LoadPANKDTree()

	if err != nil {
		fmt.Println("Error rebuilding KD tree from binary:", err)
		return
	}
	pk.AddPAN(&res, ref, RALtree)

	NCStree, err = pk.LoadNCSKDTree()

	if err != nil {
		fmt.Println("Error rebuilding KD tree from binary:", err)
		return
	}
	pk.AddNCS(&res, ref, RALtree)

	// Maintainance
	//io.CleanRAL()

}
