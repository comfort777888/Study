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
	t.Insert(5)
	t.Insert(7)
	t.Printinorder()
	fmt.Println("end")
	t.remove(6)
	t.Printinorder()
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

func (t *TreeNode) Printinorder(){
	if t==nil {
		return
	}
	t.left.Printinorder()
	fmt.Println(t.val)
	t.right.Printinorder()
}


func (t *TreeNode) remove(value int) *TreeNode {

	if t == nil {
		return nil
	}

	if value < t.val {
		t.left = t.left.remove(value)
		return t
	}
	if value > t.val {
		t.right = t.right.remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)
	return t
}

