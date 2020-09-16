func numIdenticalPairs(nums []int) int {
    length := len(nums)
    count :=0
    for i:=0;i<length-1;i++{
        for j:=i+1;j<length;j++{
            if nums[i]==nums[j]{
                count++
            }
        }
    }
    return count
}