func invertTree(root *TreeNode) *TreeNode {
    if root==nil{
        return root
    }else{
        var temp *TreeNode
        temp=root.Right
        root.Right=root.Left
        root.Left=temp
        root.Right=invertTree(root.Right)
        root.Left=invertTree(root.Left)
        return root
    }
}