package binarytree

import (
	"fmt"
)

type Branch struct {
	left  *Branch
	right *Branch
	data  int64
}

type Tree struct {
	Root *Branch
}

func CreateBinaryTree(treeValue []int64, tree *Tree) {
	for _, objTreevalue := range treeValue {
		tree.insert(objTreevalue)
	}
}

func (objTree *Tree) insert(data int64) *Tree {
	if objTree.Root == nil {
		objTree.Root = &Branch{data: data, left: nil, right: nil}
	} else {
		objTree.Root.Insert(data)
	}
	return objTree
}

func (objBranch *Branch) Insert(data int64) {
	if objBranch == nil {
		return
	} else if data <= objBranch.data {
		if objBranch.left == nil {
			objBranch.left = &Branch{data: data, left: nil, right: nil}
		} else {
			objBranch.left.Insert(data)
		}
	} else {
		if objBranch.right == nil {
			objBranch.right = &Branch{data: data, left: nil, right: nil}
		} else {
			objBranch.right.Insert(data)
		}
	}
}

//------for in order
func (objTree *Tree) InOrder() {
	var treeData []int64
	if objTree.Root == nil {
		fmt.Println("Tree is empty")
		return
	} else {
		treeData = objTree.Root.InOrderBranchTraverse(treeData)
	}
	fmt.Println(treeData)
}

func (objBranch *Branch) InOrderBranchTraverse(treeData []int64) []int64 {
	if objBranch.left != nil {
		treeData = objBranch.left.InOrderBranchTraverse(treeData)
	}
	treeData = append(treeData, objBranch.data)
	if objBranch.right != nil {
		treeData = objBranch.right.InOrderBranchTraverse(treeData)
	}

	return treeData
}

//------for pre order
func (objTree *Tree) PreOrder() {
	var treeData []int64
	if objTree.Root == nil {
		fmt.Println("Tree is empty")
		return
	} else {
		treeData = objTree.Root.PreOrderBranchTraverse(treeData)
	}
	fmt.Println(treeData)
}

func (objBranch *Branch) PreOrderBranchTraverse(treeData []int64) []int64 {
	treeData = append(treeData, objBranch.data)
	if objBranch.left != nil {
		treeData = objBranch.left.PreOrderBranchTraverse(treeData)
	}

	if objBranch.right != nil {
		treeData = objBranch.right.PreOrderBranchTraverse(treeData)
	}
	return treeData
}

//------for post order
func (objTree *Tree) PostOrder() {
	var treeData []int64
	if objTree.Root == nil {
		fmt.Println("Tree is empty")
		return
	} else {
		treeData = objTree.Root.PostOrderBranchTraverse(treeData)
	}
	fmt.Println(treeData)
}

func (objBranch *Branch) PostOrderBranchTraverse(treeData []int64) []int64 {

	if objBranch.left != nil {
		treeData = objBranch.left.PostOrderBranchTraverse(treeData)
	}

	if objBranch.right != nil {
		treeData = objBranch.right.PostOrderBranchTraverse(treeData)
	}
	treeData = append(treeData, objBranch.data)
	return treeData
}
