func minDiffInBST(root *TreeNode) int {
    var findmax func(root *TreeNode) int
    var findmin func(root *TreeNode) int
    findmax=func(root *TreeNode) int{
        if root.Right!=nil{
            return findmax(root.Right)
        }
        return root.Val
    }
    findmin=func(root *TreeNode) int{
        if root.Left!=nil{
            return findmin(root.Left)
        }
        return root.Val
    }
    if root==nil{
        return -1
    }else{
        nr,nl,sr,sl:=-1,-1,-1,-1
        if root.Right!=nil{
            nr=findmin(root.Right)-root.Val
            sr=minDiffInBST(root.Right)
        }
        if root.Left!=nil{
            nl=root.Val-findmax(root.Left)
            sl=minDiffInBST(root.Left)
        }
        min:=-1
        if nr!=-1{
            if min==-1{
                min=nr
            }else{
                if min>nr{
                    min=nr
                }
            }
        }
        if nl!=-1{
            if min==-1{
                min=nl
            }else{
                if min>nl{
                    min=nl
                }
            }
        }
        if sr!=-1{
            if min==-1{
                min=sr
            }else{
                if min>sr{
                    min=sr
                }
            }
        }
        if sl!=-1{
            if min==-1{
                min=sl
            }else{
                if min>sl{
                    min=sl
                }
            }
        }
        return min 
    }
}