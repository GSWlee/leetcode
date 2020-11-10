func nextPermutation(nums []int)  {
    length:=len(nums)
    i:=length-2
    for ;i>=0;i--{
        if nums[i]<nums[i+1]{
            break
        }
    }
    if i<0{
        for j:=length;j>i;j--{
            for m:=i+1;m<j-1;m++{
                if nums[m]>nums[m+1]{
                    temp:=nums[m]
                    nums[m]=nums[m+1]
                    nums[m+1]=temp
                }
            }
        }
        return  
    }
    min:=i+1
    for j:=length-1;j>i;j--{
        if nums[j]>nums[i]&&nums[j]<=nums[min]{
            min=j
        }
    }
    temp:=nums[min]
    nums[min]=nums[i]
    nums[i]=temp
    for j:=length;j>i;j--{
        for m:=i+1;m<j-1;m++{
            if nums[m]>nums[m+1]{
                temp:=nums[m]
                nums[m]=nums[m+1]
                nums[m+1]=temp
            }
        }
    }
}