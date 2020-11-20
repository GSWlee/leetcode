
 func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    fakehead:=&ListNode{Next:head}
    lastsort,curr:=head,head.Next
    for curr!=nil{
        if lastsort.Val<=curr.Val{
            lastsort=lastsort.Next
        }else{
            pre:=fakehead
            for pre!=nil{
                if pre.Next.Val<=curr.Val{
                    pre=pre.Next
                }else{
                    break
                }
            }
            lastsort.Next=curr.Next
            temp:=pre.Next
            pre.Next=curr
            curr.Next=temp
        }
        curr=lastsort.Next
    }
    return fakehead.Next
}