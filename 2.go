func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
    flag:=0
    temp:=&ListNode{}
    for l1!=nil||l2!=nil||flag==1{
        n1,n2:=0,0
        if l1!=nil{
            n1=l1.Val
            l1=l1.Next
        }
        if l2!=nil{
            n2=l2.Val
            l2=l2.Next
        }
        num:=n1+n2+flag
        if num>=10{
            flag=1
            num=num-10
        }else{
            flag=0
        }
        if head==nil{
            head=&ListNode{Val: num}
            temp=head
        }else{
            temp.Next=&ListNode{Val:num}
            temp=temp.Next
        }
    }
    return
}