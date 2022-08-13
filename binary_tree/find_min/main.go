package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	t :=&TreeNode{val: 8}
	t.Insert(4)
	t.Insert(2)
	t.Insert(3)
	t.Insert(10)
	t.Insert(6)
	t.Insert(7)
	min:=t.FindMin()
	fmt.Println(min)
	}

func (t *TreeNode) Insert(value int) error{
	if t == nil {
		return errors.New("tree doesn't exist")
	}

	if t.val == value {
		return errors.New("value already exists")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func (t *TreeNode) PrintInorder() {
	if t==nil {
		return
	}
	t.left.PrintInorder()
	fmt.Println(t.val)
	t.right.PrintInorder()
}

func (t *TreeNode) FindMin() int {
	if t.left!=nil {
		return t.left.FindMin()
	} else {
		return t.val
	}	
}