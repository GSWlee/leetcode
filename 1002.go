func commonChars(A []string) []string {
    ans :=[]string{}
    length:=len(A)
    flag1:=true
    var find func(str string,target rune)bool
    var delete func(str string,target rune)string
    find=func(str string,target rune)bool{
        for _,i:=range(str){
            if i==target{
                return true
            }
        }
        return false
    }
    delete=func(str string,target rune)string{
        i:=0
        if len(str)==1{
            flag1=false
        }
        for ;i<len(str);i++{
            if target==rune(str[i]){
                if  i==len(str){
                    return str[:i]
                }
                return str[:i]+str[i+1:]
            }
        }
        return str
    }
    for _,t:=range(A[0]){
        flag:=true
        for i:=1;i<length;i++{
            if find(A[i],rune(t)){
                A[i]=delete(A[i],rune(t))
            }else{
                flag=false
                break
            }
        }
        if flag{
            ans=append(ans,string(t))
        }
        if flag1==false{
            return ans
        }
    }
    return ans
}