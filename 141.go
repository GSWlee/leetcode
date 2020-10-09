func hasCycle(head *ListNode) bool {
    fast,slow:=head,head
    if fast==nil{
        return false
    }else{
        fast=fast.Next
    }
    for fast!=slow{
        if fast==nil{
            return false
        }else{
            fast=fast.Next
        }
        if fast==nil{
            return false
        }else{
            fast=fast.Next
        }
        if slow==nil{
            return false
        }else{
            slow=slow.Next
        }
    }
    return true
}