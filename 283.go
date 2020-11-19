func moveZeroes(nums []int)  {
    flag:=-1
    for i:=0;i<len(nums);i++{
        if nums[i]==0{
            if flag==-1{
                flag=i
            }
        }else{
            if flag==-1{
                continue
            }else{
                nums[flag]=nums[i]
                flag++
                nums[i]=0
            }
        }
    }
}