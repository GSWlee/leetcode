func oddEvenList(head *ListNode) *ListNode {
    if head==nil{
        return head
    }
    evens:=head.Next
    if evens==nil{
        return head
    }
    tempo,tempe:=head,evens
    for tempo!=nil&&tempe!=nil{
        if tempe!=nil{
            tempo.Next=tempe.Next
            tempo=tempo.Next
        }
        if tempo!=nil{
            tempe.Next=tempo.Next
            tempe=tempe.Next
        }
    }
    tempo=head
    for tempo.Next!=nil{
        tempo=tempo.Next
    }
    tempo.Next=evens
    return head
}