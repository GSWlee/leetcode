func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left==right{
        return head
    }
    start,end:=head,head
    for i:=1;i<left;i++{
        start=start.Next
    }
    for i:=1;i<right;i++{
        end=end.Next
    }
    tail:=start
    last:=end.Next
    temp:=&ListNode{}
    for start!=last{
        t:=start.Next
        start.Next=temp
        temp=start
        start=t
    }
    if left==1{
        head=temp
        tail.Next=last
    }else{
        begin:=head
        for i:=2;i<left;i++{
            begin=begin.Next
        }
        begin.Next=temp
        tail.Next=last
    }
    return head
}