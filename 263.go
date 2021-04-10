func isUgly(n int) bool {
    if n<=0{
        return false
    }
    nums:=[]int{2,3,5}
    for n!=1{
        change:=false
        for i:=0;i<3;i++{
            if n%nums[i]==0{
                n=n/nums[i]
                change=true
                break
            }
        }
        if !change{
            return false
        }
    }
    return true
}