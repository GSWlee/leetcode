func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root==nil{
        root=&TreeNode{val,nil,nil}
        return root
    }else{
        if val>root.Val{
            root.Right=insertIntoBST(root.Right,val)
        }else{
            root.Left=insertIntoBST(root.Left,val)
        }
        return root
    }
}