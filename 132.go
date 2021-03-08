func minCut(s string) int {
    var isPalind func(s string) bool
    isPalind=func(s string) bool{
        i,j:=0,len(s)-1
        for i<j{
            if s[i]!=s[j]{
                return false
            }
            i++
            j--
        }
        return true
    }
    dp :=make([]int,len(s)+1)
    for i:=1;i<=len(s);i++{
        min:=2000
        for j:=0;j<i;j++{
            if isPalind(s[j:i]){
                temp:=0
                if j!=0{
                    temp=1+dp[j]
                }
                if temp<min{
                    min=temp
                }
            }
        }
        dp[i]=min
    }
    return dp[len(dp)-1]
}