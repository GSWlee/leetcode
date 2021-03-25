func deleteDuplicates(head *ListNode) *ListNode {
    new:=&ListNode{0,nil}
    ans:=new
    for head!=nil&&head.Next!=nil{
        if head.Val!=head.Next.Val{
            new.Next=head
            new=new.Next
            head=head.Next
        }else{
            temp:=head.Val
            for head!=nil&&head.Val==temp{
                head=head.Next
            }
        }
    }
    if head==nil{
        new.Next=nil
        return ans.Next
    }else{
        new.Next=head
        new.Next.Next=nil
        return ans.Next
    }
}