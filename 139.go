func wordBreak(s string, wordDict []string) bool {
	dp:=make([]bool,len(s)+1)
    for i:=0;i<=len(s);i++{
        for _,j:=range(wordDict){
            if i==len(j){
                if s[:i]==j{
                    dp[i]=true
                }
            }else{
                t:=i-len(j)
                if t>0{
                    if s[t:i]==j&&dp[t]{
                        dp[i]=true
                    }
                }
            }
        }
    }
	return dp[len(s)]
}