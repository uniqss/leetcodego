package main

import "log"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var null = 0

func main() {
	var x int
	var y int
	nums := []int{1,2,3,null,4,null,5}
	x = 5
	y = 4
	var root []*TreeNode
	for i:=range nums{
		if i != null {
			node := &TreeNode{nums[i], nil, nil}
			root = append(root, node)
		}
	}
	for i:=0;i< len(nums);i++ {
		if i > 0 {
			root[i].Left = root[i - 1]
		}
		if i < len(nums) - 1 {
			root[i].Right = root[i + 1]
		}
	}
	result := isCousins(root[0], x, y)
	log.Println(result)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	var nodes []*TreeNode
	return false
}
