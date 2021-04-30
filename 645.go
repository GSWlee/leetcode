func findErrorNums(nums []int) []int {
    sum,one,flag:=0,0,true
    sort.Ints(nums)
    for _,v:=range nums{
        if flag{
            if v==one{
                flag=false
            }else{
                one=v
            }
        }
        sum+=v
    }
    miss:=(1+len(nums))*len(nums)/2+one-sum
    return append(append([]int{},one),miss)
}