//递归解法
func swapPairs(head *ListNode) *ListNode {
    if head!=nil&&head.Next!=nil{
        first,temp:=head,head.Next.Next
        head=head.Next
        head.Next=first
        head.Next.Next=swapPairs(temp)
        return head
    }else{
        return head
    }
}

//迭代解法
func swapPairs(head *ListNode) *ListNode {
    headhead:=&ListNode{0,head}
    temp:=headhead
    for temp.Next!=nil&&temp.Next.Next!=nil{
        first,last:=temp.Next,temp.Next.Next
        temp.Next=last
        first.Next=last.Next
        last.Next=first
        temp=first
    }
    return headhead.Next
}