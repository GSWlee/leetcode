func getleaf(root *TreeNode) []int{
    leaves:=[]int{}
    var add func(root *TreeNode)
    add=func(root *TreeNode){
        if root==nil{
            return
        }
        if root.Left==nil&&root.Right==nil{
            leaves=append(leaves,root.Val)
        }
        add(root.Left)
        add(root.Right)
        return
    }
    add(root)
    return leaves
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    tree1:=getleaf(root1)
    tree2:=getleaf(root2)
    if len(tree1)==len(tree2){
        for i:=range(tree1){
            if tree1[i]!=tree2[i]{
                return false
            }
        }
        return true
    }
    return false
}