func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func getdeep(root *TreeNode) int{
    switch{
        case root==nil:
            return 0
        case root.Left==nil&&root.Right==nil:
            return 1
        default:
            q := Max(getdeep(root.Right),getdeep(root.Left))
            return q+1
    }
}

func addnum(root *TreeNode,i int, arr[][]int){
    if root==nil{
        return
    } else{
        addnum(root.Left,i+1,arr)
        arr[i]=append(arr[i],root.Val)
        addnum(root.Right,i+1,arr)
    }
}

func averageOfLevels(root *TreeNode) []float64 {
    k :=getdeep(root)
    var arr [][]int
    for i:=0;i<k;i++{
        var temp []int
        arr=append(arr,temp)
    }
    addnum(root,0,arr)
    ans := make([]float64,k)
    for i:=0;i<k;i++{
        sum := 0.0
        count :=0
        for _,value := range arr[i]{
            sum =sum+float64(value)
            count++
        }
        ans[i]=sum/float64(count)
    }
    return ans
}