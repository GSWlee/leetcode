func removeDuplicates(nums []int) int {
    length:=len(nums)
    if length==0{
        return 0
    }
    temp,i:=0,0
    num,now:=0,nums[0]
    for;i<length;{
        if now==nums[i]{
            num++
        }else{
            num=1
            now=nums[i]
        }
        switch num{
            case 1:
                nums[temp]=nums[i]
                temp++
                i++
            case 2:
                nums[temp]=nums[i]
                temp++
                i++
            default:
                i++
        }
    }
    return temp
}