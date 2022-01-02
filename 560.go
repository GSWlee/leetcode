//前缀和 n^2
func subarraySum(nums []int, k int) int {
    ans:=make([]int,len(nums))
    R:=0
    for k,v:=range nums{
        R+=v
        ans[k]=R
    }
    temp:=0
    for j:=len(ans)-1;j>=0;j--{
        if ans[j]==k{
            temp++
        }
        for i:=0;i<j;i++{
            if ans[j]-ans[i]==k{
                temp++
            }
        }
	}
}

//前缀和加哈希
func subarraySum(nums []int, k int) int {
    count, pre := 0, 0
    m := map[int]int{}
    m[0] = 1
    for i := 0; i < len(nums); i++ {
        pre += nums[i]
        if _, ok := m[pre - k]; ok {
            count += m[pre - k]
        }
        m[pre] += 1
    }
    return count
}
    