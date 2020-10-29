func sumNumbers(root *TreeNode) int {
    var nums=[]int{}
    var add func(root *TreeNode,value int)
    add=func(root *TreeNode,value int){
        if root==nil{
            return
        }
        if root.Left==nil&&root.Right==nil{
            nums=append(nums,value+root.Val)
            return
        }else{
            temp:=value+root.Val
            if root.Left!=nil{
                add(root.Left,temp*10)
            }
            if root.Right!=nil{
                add(root.Right,temp*10)
            }
        }
    }
    add(root,0)
    sum:=0
    for _,j:=range(nums){
        sum+=j
    }
    return sum
}