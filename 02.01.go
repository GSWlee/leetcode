// 无脑版本
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

// 位运算
func removeDuplicateNodes(head *ListNode) *ListNode {
    point:=head
    father:=head
    if point==nil{
        return head
    }
    var a [630]int
    for i := range(a){
        a[i]=0;
    }
    var update func(a [630]int,i,j int)[630]int
    update=func(a [630]int,i,j int) [630]int{
        a[i]=a[i]+int(math.Pow(2.0,float64(j)))
        return a
    }
    
    var find func(a [630]int,target int)bool
    find=func(a [630]int,target int) bool{
        i,j:=target/32,target%32
        if a[i]&int(math.Pow(2.0,float64(j)))!=0{
            return true
        }
        return false
    }
    a=update(a,point.Val/32,point.Val%32)
    point=point.Next
    for point!=nil{
        if find(a,point.Val){
            father.Next=point.Next
            point=point.Next
        }else{
            a=update(a,point.Val/32,point.Val%32)
            point=point.Next
            father=father.Next
        }
    }
    return head
}