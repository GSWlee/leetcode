func increasingTriplet(nums []int) bool {
    small,mid:=int(^uint(0) >> 1),int(^uint(0) >> 1)
    if len(nums)<3{
        return false
    }else{
        for _,v:=range nums{
            if v<=small{
                small=v
            }else if v<=mid{
                mid=v
            }else{
                return true
            }
        }
    }
    return false
}