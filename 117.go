type peer struct{
	Val *Node
	deep int
}

func connect(root *Node) *Node {
   var nodes[]peer
   if root==nil{
	   return root
   }else{
	   nodes=append(nodes,peer{root,0})
	   i:=0
	   for len(nodes)>i{
		   if nodes[i].Val.Left!=nil{
			   nodes=append(nodes,peer{nodes[i].Val.Left,nodes[i].deep+1})
		   }
		   if nodes[i].Val.Right!=nil{
			   nodes=append(nodes,peer{nodes[i].Val.Right,nodes[i].deep+1})
		   }
		   i++
	   }
	   for i=0;i<len(nodes)-1;i++{
		   if nodes[i].deep==nodes[i+1].deep{
			   nodes[i].Val.Next=nodes[i+1].Val
		   }else{
			   nodes[i].Val.Next=nil
		   }
	   }
	   nodes[i].Val.Next=nil
	   return root
   }
}


// 方法二，降低时间复杂度
func connect(root *Node) *Node {
    if root==nil{
        return rootß
    }else{
        q:=[]*Node{root}
        for len(q)>0{
            temp:=q
            q=nil
            for i,_:=range(temp){
                if i+1<len(temp){
                    temp[i].Next=temp[i+1]
                }
                if temp[i].Left!=nil{
                    q=append(q,temp[i].Left)
                }
                if temp[i].Right!=nil{
                    q=append(q,temp[i].Right)
                }
            }
        }
        return root
    }
}