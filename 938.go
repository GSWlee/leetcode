func rangeSumBST(root *TreeNode, low int, high int) int {
    if root==nil{
        return 0
    }else if root.Val>high{
        return rangeSumBST(root.Left,low,high)
    }else if root.Val<low{
        return rangeSumBST(root.Right,low,high)
    }else{
        return root.Val+rangeSumBST(root.Left,low,high)+rangeSumBST(root.Right,low,high)
    }
}