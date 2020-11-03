func validMountainArray(A []int) bool {
    start,end:=1,len(A)
    if end<3{
        return false
    }
    flag:=true
    top:=0
    for ;start<end;start++{
        if flag{
            if A[start]==A[start-1]{
                return false
            }else if A[start]<A[start-1]{
                flag=false
                top=start-1
            }else{
                continue
            }
        }else{
            if A[start]>=A[start-1]{
                return false
            }
        }
       
    }
    if top!=0{
        return true
    }else{
        return false
    }
    return true
}