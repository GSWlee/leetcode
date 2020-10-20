func findmid(head *ListNode)*ListNode{
    fast,slow:=head,head
    for fast.Next!=nil&&fast.Next.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
    }
    return slow
}
func reverse(head *ListNode) *ListNode{
    var ans,cur *ListNode = nil,head
    for cur!=nil{
        temp:=ans
        ans=cur
        cur=cur.Next
        ans.Next=temp
    }
    return ans
}
func merge(l1,l2 *ListNode){
    cur:=l1
    for l2!=nil{
        temp:=cur.Next
        cur.Next=l2
        l2=l2.Next
        cur.Next.Next=temp
        cur=cur.Next.Next
    }
}

func reorderList(head *ListNode)  {
    if head==nil{
        return
    }
    l1,l2:=head,findmid(head)
    temp:=l2
    l2=l2.Next
    temp.Next=nil
    l2=reverse(l2)
    merge(l1,l2)
}