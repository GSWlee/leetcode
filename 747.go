func dominantIndex(nums []int) int {
    max,taget:=0,0
    for i:=1;i<len(nums);i++{
        if taget==-1{
            if nums[i]>=2*nums[max]{
                taget=i
            }
        }else{
            if nums[taget]<2*nums[i]{
                if nums[i]>=2*nums[taget]{
                    taget=i
                }else{
                    if nums[taget]>nums[i]{
                        max=taget
                    }else{
                        max=i
                    }
                    taget=-1
                }    
            }
        }
    }
    return taget
}