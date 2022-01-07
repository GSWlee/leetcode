func maxDepth(s string) int {
    ans:=0
    temp:=0
    stack:=[]rune{}
    for _,v:=range s{
        if v=='('{
            stack=append(stack,v)
            temp++
            if temp>ans{
                ans=temp
            }
        }else if v==')'{
            stack=stack[:len(stack)-1]
            temp--
        }else{
            continue
        }
    }
    return ans
}