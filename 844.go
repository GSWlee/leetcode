func backspaceCompare(S string, T string) bool {
    var reduce func(S string)string
    reduce=func(S string)string{
        var ans string
        for i,j:=range(S){
            if j=='#'{
                if len(ans)!=0{
				    ans=ans[:len(ans)-1]
			    }
            }else{
                ans=ans+S[i:i+1]
            }
        }
        return ans
    }
    S=reduce(S)
    T=reduce(T)
    if len(S)!=len(T){
        return false
    }
    for i,_:=range(S){
        if S[i]!=T[i]{
            return false
        }
    }
    return true
}