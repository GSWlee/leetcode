func inorderTraversal(root *TreeNode) []int {
    var ans []int
    if root == nil{
        return ans
    }
    var inorder func(root* TreeNode,ans []int)[]int
    inorder=func(root* TreeNode,ans []int)[]int{
        if root==nil{
            return ans
        }
        ans=inorder(root.Left,ans)
        ans=append(ans,root.Val)
        ans=inorder(root.Right,ans)
        return ans
    }   
    ans=inorder(root,ans)
    return ans
}