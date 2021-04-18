func removeDuplicates(nums []int) int {
    length:=len(nums)
    i,j:=1,0
    for ;i<length;i++{
        if nums[i]!=nums[j]{
            j++
            nums[j]=nums[i]
        }
    }
    return j+1
}