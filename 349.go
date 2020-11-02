func intersection(nums1 []int, nums2 []int) []int {
    var nofind func(nums1,nums2 []int, target int)bool
    nofind=func(nums1,nums2 []int, target int)bool{
        flag:=false
        for _,i:=range(nums1){
            if i==target{
                flag=true
            }
        }
        for _,i:=range(nums2){
            if i==target{
                flag=false
            }
        }
        return flag
    }
    var ans=[]int{}
    for _,i:=range(nums2){
        if nofind(nums1,ans,i){
            ans=append(ans,i)
        }
    }
    return ans
}