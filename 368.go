func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    dp:=make([][]int,len(nums))
    for i:=range nums{
        max:=0
        for j:=0;j<i;j++{
            if nums[i]%nums[j]==0{
                if len(dp[j])>max{
                    dp[i]=append([]int{},dp[j]...)
                    max=len(dp[j])
                }
            }
        }
        dp[i]=append(dp[i],nums[i])
    }
    max:=0
    for i:=1;i<len(nums);i++{
        if len(dp[i])>len(dp[max]){
            max=i
        }
    }
    return dp[max]
}