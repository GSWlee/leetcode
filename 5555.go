func countVowelStrings(n int) int {
    var ans=[][]int{}
    var temp=[]int{1,1,1,1,1}
    ans=append(ans,temp)
    for i:=1;i<=n;i++{
        var t=[]int{0,0,0,0,1}
        for j:=3;j>=0;j--{
            for m:=4;m>=j;m--{
                t[j]+=ans[i-1][m]
            }
        }
        ans=append(ans,t)
    }
    return ans[n][0]
}