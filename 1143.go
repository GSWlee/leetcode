func longestCommonSubsequence(text1 string, text2 string) int {
    m,n:=len(text1),len(text2)
    dp:=make([][]int,n+1)
    for i,_:=range dp{
        dp[i]=make([]int,m+1)
    }
    for i:=1;i<=n;i++{
        for j:=1;j<=m;j++{
            if text2[i-1]==text1[j-1]{
                dp[i][j]=dp[i-1][j-1]+1
            }else{
                if dp[i-1][j]>dp[i][j-1]{
                    dp[i][j]=dp[i-1][j]
                }else{
                    dp[i][j]=dp[i][j-1]
                }
            }
        }
    }
    return dp[n][m]
}