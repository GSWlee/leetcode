func search(nums []int, target int) int {
    length:=len(nums)
    start,end:=0,length-1
    flag:=0
    for start<=end{
        t:=(start+end)/2
        if t==length-1{
            break
        }else{
            if nums[t]>nums[t+1]{
                flag=t+1
                break
            }else{
                if nums[t]>=nums[start]{
                    start=t+1
                }else{
                    end=t-1
                }
            }
        }
    }
    start,end=0,length-1
    ans:=-1
    for start<=end{
        t:=(start+end)/2
        if nums[(t+flag)%length]==target{
            ans=(t+flag)%length
            break
        }
        if nums[(t+flag)%length]>target{
            end=t-1
        }else{
            start=t+1
        }
    }
    return ans
}