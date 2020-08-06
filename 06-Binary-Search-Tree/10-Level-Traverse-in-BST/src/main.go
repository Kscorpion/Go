package main

import "Play-with-Data-Structures/06-Binary-Search-Tree/10-Level-Traverse-in-BST/src/BST"

func main() {
	bst := BST.Constructor()
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
