func mergeTwoLists(list1 *ListNode, list2 *ListNode) (head *ListNode) {
    temp:=&ListNode{}
    for list1!=nil||list2!=nil{
        if list1==nil{
            if head==nil{
                head=&ListNode{Val: list2.Val}
                temp=head
            }else{
                temp.Next=&ListNode{Val: list2.Val}
                temp=temp.Next
            }
            list2=list2.Next
            continue
        }
        if list2==nil{
            if head==nil{
                head=&ListNode{Val: list1.Val}
                temp=head
            }else{
                temp.Next=&ListNode{Val: list1.Val}
                temp=temp.Next
            }
            list1=list1.Next
            continue
        }
        if list1.Val<list2.Val{
            if head==nil{
                head=&ListNode{Val: list1.Val}
                temp=head
            }else{
                temp.Next=&ListNode{Val: list1.Val}
                temp=temp.Next
            }
            list1=list1.Next
            continue
        }else{
            if head==nil{
                head=&ListNode{Val: list2.Val}
                temp=head
            }else{
                temp.Next=&ListNode{Val: list2.Val}
                temp=temp.Next
            }
            list2=list2.Next
            continue
        }
    }
    return
}