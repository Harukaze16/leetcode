package main

import "fmt"

 
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    var res []int
    cur := root
    for cur != nil {
        if cur.Left == nil { 
            //左孩子空，必定是首次遇到
            res = append(res,cur.Val)
            cur = cur.Right
            continue
        }
        leftMost := cur.Left
        for leftMost.Right!=nil && leftMost.Right!=cur { 
            leftMost = leftMost.Right
        }
        if leftMost.Right != nil {
            leftMost.Right = nil
            res = append(res,cur.Val)
            cur = cur.Right
        } else {
            leftMost.Right = cur
            cur = cur.Left
        }

    }
    return res
}

func main() {
    son := TreeNode{2,nil,nil}
    parent := TreeNode{1,&son,nil}
    inorderTraversal(&parent)
    fmt.Println("ok")
}