func postorderTraversal(root *TreeNode) []int {
    ans := []int{}
    var postorder func(root *TreeNode)
    postorder=func(root *TreeNode){
        if root==nil{
            return
        }else{
            postorder(root.Left)
            postorder(root.Right)
            ans=append(ans,root.Val)
        }
    }
    postorder(root)
    return ans
}