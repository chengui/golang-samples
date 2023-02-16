package treesort

import (
    "fmt"
)

type tree struct {
    value int
    left, right *tree
}

func Sort(arr []int) {
    var root *tree
    for _, v := range arr {
        root = add(root, v)
    }
    arr = arr[:0]
    var inorder func(*tree)
    inorder = func(root *tree) {
        if root != nil {
            inorder(root.left)
            arr = append(arr, root.value)
            inorder(root.right)
        }
    }
    inorder(root)
    //arr = inorder(arr[:0], root)
    fmt.Println("sort: ", arr)
}

func add(root *tree, val int) *tree {
    if root == nil {
        return &tree{value: val}
    }
    if val < root.value {
        root.left = add(root.left, val)
    } else {
        root.right = add(root.right, val)
    }
    return root
}

func inorder(arr []int, root *tree) []int {
    if root != nil {
        arr = inorder(arr, root.left)
        arr = append(arr, root.value)
        arr = inorder(arr, root.right)
    }
    return arr
}
