func buildTree(inorder []int, postorder []int) *TreeNode {
    var temp TreeNode
    if len(inorder)==0{
        return nil
    }
    if len(inorder)==1{
        temp.Val=inorder[0]
        temp.Right=nil
        temp.Left=nil
        root:=&temp
        return root
    }else{
        temp.Val=postorder[len(postorder)-1]
        var flag int
        for i :=range(inorder){
            if inorder[i]==temp.Val{
                flag=i
            }
        }
        temp.Left=buildTree(inorder[0:flag],postorder[0:flag])
        if (flag+1)==len(inorder){
            temp.Right=nil
        }else{
            temp.Right=buildTree(inorder[flag+1:],postorder[flag:len(postorder)-1])
        }
        root:=&temp
        return root
    }
    
}