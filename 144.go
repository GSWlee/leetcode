func preorderTraversal(root *TreeNode) []int {
    var ans =[]int{}
    var preorder func(root *TreeNode)
    preorder=func(root *TreeNode){
        if root==nil{
            return
        }else{
            ans=append(ans,root.Val)
            preorder(root.Left)
            preorder(root.Right)    
        }
    }
    preorder(root)
    return ans
}