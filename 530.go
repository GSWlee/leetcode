func getMinimumDifference(root *TreeNode) int {
    nums:=[]int{}
    var find func(root *TreeNode)
    find=func(root *TreeNode){
        if root==nil{
            return
        }else{
            find(root.Right)
            nums=append(nums,root.Val)
            find(root.Left)
        }
    }
    find(root)
    length:=len(nums)
    if length<2{
        return int(^uint(0) >> 1)
    }else{
        min:=int(^uint(0) >> 1)
        for i:=1;i<length;i++{
            temp:=nums[i]-nums[i-1]
            if temp<0{
                temp=-temp
            }
            if temp<min{
                min=temp
            }
        }
        return min
    }
}