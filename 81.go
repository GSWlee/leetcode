func search(nums []int, target int) bool {
    length:=len(nums)
	start,end:=0,length-1
    var searchA func(start int,end int) bool
    searchA=func(start int,end int) bool{
        if start>end{
            return false
        }else{
            t:=(start+end)/2
		    if nums[t]==target{
			    return true
		    }else{
                if nums[start]<nums[t]{
                    if nums[t]>target&&target>=nums[start]{
                        return searchA(start,t-1)
                    }else{
                        return searchA(t+1,end)
                    }
                }else if nums[end]>nums[t]{
                    if nums[t]<target&&nums[end]>=target{
                        return searchA(t+1,end)
                    }else{
                        return searchA(start,t-1)
                    }
                }else{
                    return searchA(start,t-1)||searchA(t+1,end)
                }
            }
        }
    }
	return searchA(start,end)
}