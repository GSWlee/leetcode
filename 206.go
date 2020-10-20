func reverseList(head *ListNode) *ListNode {
    var ans,cur *ListNode = nil,head
    for cur!=nil{
        temp:=ans
        ans=cur
        cur=cur.Next
        ans.Next=temp
    }
    return ans
}