func sortArrayByParityII(A []int) []int {
    var odds,evens,ans=[]int{},[]int{},[]int{}
    for i:=0;i<len(A);i++{
        if A[i]%2==0{
            evens=append(evens,A[i])
        }else{
            odds=append(odds,A[i])
        }
    }
    j:=0
    for i:=0;i<len(A);i++{
        if i%2==0{
            ans=append(ans,evens[j])
        }else{
            ans=append(ans,odds[j])
            j++
        }
    }
    return ans
}