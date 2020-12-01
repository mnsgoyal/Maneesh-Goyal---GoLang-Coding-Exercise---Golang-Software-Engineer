package main

import (
	"assignment/binarytree"
	"assignment/houserobber"
	"fmt"
)

func main() {
	fmt.Println("Binary Tree Assignment------------START")
	//------zero value of array will be root of tree
	tree := binarytree.Tree{}
	treeValue := []int64{100, 20, 150, 10, 5, 25, 102, 456}
	binarytree.CreateBinaryTree(treeValue, &tree)

	//-------Insert Operation on tree
	var insertValue int64 = 54
	fmt.Println("Inserting the node in binary tree with value ", insertValue)
	tree.Root.Insert(insertValue)

	//--------InOrder Traversal
	fmt.Println("Tree value using InOrder Traversal")
	tree.InOrder()

	//--------PreOrder Traversal
	fmt.Println("Tree value using PreOrder Traversal")
	tree.PreOrder()

	//--------PostOrder Traversal
	fmt.Println("Tree value using PostOrder Traversal")
	tree.PostOrder()

	fmt.Println("Binary Tree Assignment------------END")

	fmt.Println("House Robber------------START")
	houseMoney := []int{5, 5, 3, 1}
	fmt.Println("Input value for house Robbery :", houseMoney)
	houserobber.HouseRobber(houseMoney)
	fmt.Println("House Robber------------END")
}
