func canCross(stones []int) bool {
    n:=len(stones)
    dp:=make([][]bool,n)
    for i:=1;i<n;i++{
        if stones[i]-stones[i-1]>i{
            return false
        }
    }
    for i:=0;i<n;i++{
        dp[i]=make([]bool,n)
    }
    dp[0][0]=true
    for i:=1;i<n;i++{
        for j:=i-1;j>=0;j--{
            if stones[i]-stones[j]>j+1{
                break
            }
            k:=stones[i]-stones[j]
            dp[i][k]=dp[j][k-1]||dp[j][k]||dp[j][k+1]
            if i==n-1&&dp[i][k]{
                return true
            }
        }
    }
    return false
}
