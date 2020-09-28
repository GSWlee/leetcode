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