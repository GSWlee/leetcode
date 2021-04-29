func max(a int,b int) int{
    if a<b{
        return b
    }
    return a
}

func maxProfit(prices []int) int {
    length:=len(prices)
    dp:=make([][5]int,length)
    dp[0][1]=-prices[0]
    dp[0][2],dp[0][3]=int(math.NaN()),int(math.NaN())
    for i:=1;i<length;i++{
        dp[i][1]=max(dp[i-1][1],-prices[i])
        dp[i][2]=max(dp[i-1][2],dp[i-1][1]+prices[i])
        if dp[i-1][2]!=int(math.NaN()){
            if dp[i-1][3]!=int(math.NaN()){
                dp[i][3]=max(dp[i-1][2]-prices[i],dp[i-1][3])
            }else{
                dp[i][3]=dp[i-1][2]-prices[i]
            }
        }else{
            dp[i][3]=int(math.NaN())
        }
        if dp[i-1][3]!=int(math.NaN()){
            if dp[i-1][4]!=int(math.NaN()){
                dp[i][4]=max(dp[i-1][3]+prices[i],dp[i-1][4])
            }else{
                dp[i][4]=dp[i-1][3]+prices[i]
            }
        }else{
            dp[i][4]=int(math.NaN())
        }

    }
    MAX:=0
    for i:=1;i<5;i++{
        if dp[length-1][i]>MAX{
            MAX=dp[length-1][i]
        }
    }
    return MAX
}