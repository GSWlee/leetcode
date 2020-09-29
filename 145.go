//递归
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
//非递归
func postorderTraversal(root *TreeNode) []int {
    ans:=[]int{}
    if root==nil{
        return ans
    }else{
        temp:=[]*TreeNode{root}
        for len(temp)>0{
            q:=temp[len(temp)-1]
            if q.Left!= nil{
                temp=append(temp,q.Left)
                q.Left=nil
            }else if q.Right!=nil{
                temp=append(temp,q.Right)
                q.Right=nil
            }else{
                ans=append(ans,q.Val)
                temp=temp[:len(temp)-1]
            }
        }
        return ans
    }
}