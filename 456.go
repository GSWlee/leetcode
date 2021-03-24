func find132pattern(nums []int) bool {
    for i:=1;i<len(nums)-1;i++{
        min:=nums[i]
        for n:=0;n<i;n++{
            if nums[n]<min{
                min=nums[n]
            }
        }
        if min==nums[i]{
            continue
        }
        for n:=i+1;n<len(nums);n++{
            if nums[n]<nums[i]&&nums[n]>min{
                return true
            }
        }
    }
    return false
}