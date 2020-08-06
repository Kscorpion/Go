package main

import (
	"Play-with-Data-Structures/12-AVL-Tree/07-Remove-Elements-in-AVL-Tree/src/AVLTree"
	"Play-with-Data-Structures/12-AVL-Tree/07-Remove-Elements-in-AVL-Tree/src/BSTMap"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Pride and Prejudice")

	filename, _ := filepath.Abs("12-AVL-Tree/06-LR-and-Rl/pride-and-prejudice.txt")
	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words: ", len(words))

	// Test BST
	startTime := time.Now()

	bst := BSTMap.Constructor()
	for _, word := range words {
		if bst.Contains(word) {
			bst.Set(word, bst.Get(word).(int)+1)
		} else {
			bst.Add(word, 1)
		}
	}
	for _, word := range words {
		bst.Contains(word)
	}

	diffTime := time.Now().Sub(startTime)
	fmt.Println("BST: ", diffTime)

	// Test AVL Tree
	startTime = time.Now()

	avl := AVLTree.Constructor()
	for _, word := range words {
		if avl.Contains(word) {
			avl.Set(word, avl.Get(word).(int)+1)
		} else {
			avl.Add(word, 1)
		}
	}
	for _, word := range words {
		avl.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)
	fmt.Println("AVL: ", diffTime)
}
