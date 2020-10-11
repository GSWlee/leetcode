// 递归解法，复杂度超高
func canPartition(nums []int) bool {
    sum:=0
    var find func(nums []int, target int)bool
    find=func(nums []int, target int)bool{
        if len(nums)==0{
            return false
        }
        if nums[0]==target{
            return true
        }
        q,qq:=target-nums[0],nums[1:]
        if q<0{
            return false||find(qq,target)
        }else{
            return find(qq,target)||find(qq,q)
        }
        
    }
    for _,temp:=range(nums){
        sum+=temp
    }
    if sum%2==1{
        return false
    }else{
        return find(nums,sum/2)
    }
}

//理解成0-1背包问题（dp来做）
func canPartition(nums []int) bool {
    sum:=0
    if len(nums)==0{
        return true
    }
    for _,temp:=range(nums){
        sum+=temp
    }
    if sum%2==1{
        return false
    }else{
        target:=sum/2
        dp:=[][]bool{}
        for i,_:=range(nums){
            if i==0{
                temp:=make([]bool,target+1)
                if nums[i]<=target{
                    temp[nums[i]]=true
                }
                dp=append(dp,temp)
            }else{
                temp:=make([]bool,target+1)
                if nums[i]<=target{
                    temp[nums[i]]=true
                }
                for j:=0;j<=target;j++{    
                    if (j-nums[i])>=0{
                        temp[j]=dp[i-1][j]||dp[i-1][j-nums[i]]||temp[j]
                    }else{
                        temp[j]=dp[i-1][j]||temp[j]
                    }
                }
                dp=append(dp,temp)
            }
        }
        return dp[len(nums)-1][target]
    }
}