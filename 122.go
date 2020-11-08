func maxProfit(prices []int) int {
	length:=len(prices)
	dp:=make([][2]int,length)
    dp[0][1]=-prices[0]
	for i:=1;i<length;i++{
        if dp[i-1][0]>dp[i-1][1]+prices[i]{
            dp[i][0]=dp[i-1][0]
        }else{
            dp[i][0]=dp[i-1][1]+prices[i]
        }
        if dp[i-1][1]>dp[i-1][0]-prices[i]{
            dp[i][1]=dp[i-1][1]
        }else{
            dp[i][1]=dp[i-1][0]-prices[i]
        }
	}
    if dp[length-1][0]>dp[length-1][1]{
        return dp[length-1][0]
    }else{
        return dp[length-1][1]
    }
}