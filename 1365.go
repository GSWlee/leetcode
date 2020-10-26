func smallerNumbersThanCurrent(nums []int) []int {
    tag :=make([]int,101)
    for _,i:=range(nums){
        tag[i]++
    }
    for i:=1;i<101;i++{
        tag[i]=tag[i-1]+tag[i]
    }
    for i,_:=range(nums){
        if nums[i]==0{
            nums[i]=0
        }else{
            nums[i]=tag[nums[i]-1]
        }   
    }
    return nums
}