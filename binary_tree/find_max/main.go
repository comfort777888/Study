package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func main() {
	t:=&TreeNode{val:8}
	t.Insert(4)
	t.Insert(2)
	t.Insert(3)
	t.Insert(10)
	t.Insert(6)
	t.Insert(7)
	//t.Printinorder()
	max:=t.FindMax()
	fmt.Println(max)
}

func (t *TreeNode) Insert (new int) error {
if t==nil {
	return errors.New("tree doesn't exist")
}
if new==t.val {
	return errors.New("value already exist")
}
if new<t.val {
	if t.left==nil {
		t.left=&TreeNode{val:new}
		return nil
	}
	return t.left.Insert(new)
}

if new>t.val {
	if t.right==nil {
		t.right=&TreeNode{val:new}
		return nil
	}
	return t.right.Insert(new)
}
return nil
}

// func (t *TreeNode) Printinorder(){
// 	if t==nil {
// 		return
// 	}
// 	t.left.Printinorder()
// 	fmt.Println(t.val)
// 	t.right.Printinorder()
// }

func (t *TreeNode) FindMax() int {
	if t.right!=nil {
		return t.right.FindMax()
	} else {
		return t.val
	}
}