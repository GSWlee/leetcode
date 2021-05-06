func decode(encoded []int, first int) []int {
    ans:=make([]int,1)
    ans[0]=first
    for i:=range encoded{
        temp:=ans[i]^encoded[i]
        ans=append(ans,temp)
    }
    return ans
}