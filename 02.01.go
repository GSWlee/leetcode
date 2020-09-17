func removeDuplicateNodes(head *ListNode) *ListNode {
    point:=head
    father:=head
    if point==nil{
        return head
    }
    var a []int
    var find func(a []int,target int)bool
    find=func(a []int,target int) bool{
        for _,q:=range(a){
            if q==target{
                return true
            }
        }
        return false
    }
    a=append(a,point.Val)
    point=point.Next
    for point!=nil{
        if find(a,point.Val){
            father.Next=point.Next
            point=point.Next
        }else{
            a=append(a,point.Val)
            point=point.Next
            father=father.Next
        }
    }
    return head
}