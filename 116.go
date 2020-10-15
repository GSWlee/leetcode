func connect(root *Node) *Node {
	last,now:=root,[]*Node{}
    for last!=nil{
        for last!=nil{
            if last.Left!=nil{
                now=append(now,last.Left)
            }
            if last.Right!=nil{
                now=append(now,last.Right)
            }
            last=last.Next
        }
        if len(now)==0{
            last=nil
        }else{
            for i:=0;i<len(now)-1;i++{
                now[i].Next=now[i+1]
            }
            last=now[0]
            now=[]*Node{}
        }
    } 
    return root
}