func roba(nums []int) int {
    len:=len(nums)
    dp:=make([][]int,len)
    if len==0{
        return 0
    }else if len==1{
        return nums[0]
    }else{
        dp[0]=make([]int,2)
        dp[0][0]=0
        dp[0][1]=nums[0]
        for i:=1;i<len;i++{
            dp[i]=make([]int,2)
            if dp[i-1][1]>dp[i-1][0]{
                dp[i][0]=dp[i-1][1]
            }else{
                dp[i][0]=dp[i-1][0]
            }
            dp[i][1]=nums[i]+dp[i-1][0]
        }
    }
    if dp[len-1][0]>dp[len-1][1]{
        return dp[len-1][0]
    }
    return dp[len-1][1]
}

func rob(nums []int) int{
    len:=len(nums)
    if len==0{
        return 0
    }else if len==1{
        return nums[0]
    }else{
        max1:=roba(nums[0:len-1])
        max2:=roba(nums[1:])
        if max2>max1{
            return max2
        }else{
            return max1
        }
    }
}