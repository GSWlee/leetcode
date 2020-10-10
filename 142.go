func detectCycle(head *ListNode) *ListNode {
    fast,slow:=head,head
    if fast==nil{
        return nil
    }else{
        fast=fast.Next
    }
    slow=slow.Next
    if fast==nil{
        return nil
    }else{
        fast=fast.Next
    }
    for fast!=slow{
        if fast==nil{
            return nil
        }else{
            fast=fast.Next
        }
        if fast==nil{
            return nil
        }else{
            fast=fast.Next
        }
        if slow==nil{
            return nil
        }else{
            slow=slow.Next
        }
    }
    fast=head
    for fast!=slow{
        fast=fast.Next
        slow=slow.Next
    }
    return slow
}