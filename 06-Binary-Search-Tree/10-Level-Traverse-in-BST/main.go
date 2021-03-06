package main

import "github.com/donng/Play-with-Data-Structures/06-Binary-Search-Tree/10-Level-Traverse-in-BST/BST"

func main() {
	bst := BST.New()
	nums := [...]int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}

	/////////////////
	//      5      //
	//    /   \    //
	//   3    6    //
	//  / \    \   //
	// 2  4     8  //
	/////////////////
	bst.LevelOrder()
	//bst.PreOrder()
	//fmt.Println()
	//
	//bst.PreOrderNR()
	//bst.InOrder()
	//fmt.Println()
	//
	//bst.PostOrder()
	// fmt.Println(bst)
}
