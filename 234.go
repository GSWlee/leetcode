func isPalindrome(head *ListNode) bool {
    var nums []int
    for head!=nil{
        nums=append(nums,head.Val)
        head=head.Next
    }
    first,foot:=0,len(nums)-1
    for foot>first{
        if nums[first]!=nums[foot]{
            return false
        }
        foot--
        first++
    }
    return true
}