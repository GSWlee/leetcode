func isValidSerialization(preorder string) bool {
    nums:=[]string{}
    target:=""
    for _,item := range preorder{
        if item==','{
            nums=append(nums,target)
            target=""
        }else{
            target+=string(item)
        }
    }
    nums=append(nums,target)
    ans :=2
    for i,item:=range nums{
        if ans<=0{
            return false
        }
        if i==0{
            if item=="#"{
                if len(nums)>1{
                    return false
                }
                return true
            }
            continue
        }else{
            if item=="#"{
                ans--
            }else{
                ans++
            }
        }
    }
    if ans==0{
        return true
    }
    return false
}