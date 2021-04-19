func removeElement(nums []int, val int) int {
    length:=len(nums)
    i,j:=-1,0
    for j<length{
        if nums[j]==val{
            j++
        }else{
            i++
            nums[i]=nums[j]
            j++
        }
    }
    return i+1
}