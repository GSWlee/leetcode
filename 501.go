import "strconv"
func RemoveRepeatedElement(arr []int) (newArr []int) {
    newArr = make([]int, 0)
    for i := 0; i < len(arr); i++ {
        repeat := false
        for j := i + 1; j < len(arr); j++ {
            if arr[i] == arr[j] {
                repeat = true
                break
            }
        }
        if !repeat {
            newArr = append(newArr, arr[i])
        }
    }
    return
}
func findMode(root *TreeNode) []int {
    
    nums:=make(map[string]int)
    var find func(root *TreeNode, nums map[string]int)map[string]int
    find=func(root *TreeNode, nums map[string]int)map[string]int{
        if root==nil{
            return nums
        }
        temp:=root.Val
        _,ok:=nums[strconv.Itoa(temp)]
        if ok{
            nums[strconv.Itoa(temp)]=nums[strconv.Itoa(temp)]+1
        }else{
            nums[strconv.Itoa(temp)]=1
        }
        nums=find(root.Right,nums)
        nums=find(root.Left,nums)
        return nums
    }
    nums=find(root,nums)
    ans:= []int{}
    max:=-1
    for i :=range(nums){
        if nums[i]>max{
            max=nums[i]
            q,_:=strconv.Atoi(i)
            ans=[]int{}
            ans=append(ans,q)
        }
        if nums[i]==max{
            q,_:=strconv.Atoi(i)
            ans=append(ans,q)
        }else{
            continue
        }
     
    }
    ans=RemoveRepeatedElement(ans)
    return ans
}