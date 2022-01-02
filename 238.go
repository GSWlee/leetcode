// 不使用除法
func productExceptSelf(nums []int) []int {
    first,end:=make([]int,len(nums)),make([]int,len(nums))
    first[0]=nums[0]
    for i:=1;i<len(nums);i++{
        first[i]=first[i-1]*nums[i]
    }
    end[len(nums)-1]=nums[len(nums)-1]
    for j:=len(nums)-2;j>=0;j--{
        end[j]=end[j+1]*nums[j]
    }
    ans:=make([]int,len(nums))
    for i:=0;i<len(nums);i++{
        if i==0{
            ans[i]=end[i+1]
        }else if i==len(nums)-1{
            ans[i]=first[i-1]
        }else{
            ans[i]=first[i-1]*end[i+1]
        }
    }
    return ans
}

//空间复杂度唯1
func productExceptSelf(nums []int) []int {
    ans:=make([]int,len(nums))
    ans[0]=1
    for i:=1;i<len(nums);i++{
        ans[i]=ans[i-1]*nums[i-1]
    }
    R:=1
    for j:=len(nums)-1;j>=0;j--{
        ans[j]=ans[j]*R
        R=R*nums[j]
    }
    return ans
}