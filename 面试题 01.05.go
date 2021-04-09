func oneEditAway(first string, second string) bool {
    m,n:=len(first),len(second)
    if m==n{
        change:=0
        for i:=0;i<m;i++{
            if first[i]!=second[i]{
                change++
                if change>1{
                    return false
                }
            }
        }
        return true
    }else if m-n==1{
        change:=0
        for i:=0;i<n;i++{
            if first[i]!=second[i]{
                change++
                if change>1{
                    return false
                }
                second=second[0:i]+first[i:i+1]+second[i:]
            }
        }
        if change==0||(change==1&&first[m-1]==second[m-1]){
            return true
        }
        return false
    }else if n-m==1{
        change:=0
        for i:=0;i<m;i++{
            if first[i]!=second[i]{
                change++
                if change>1{
                    return false
                }
                first=first[0:i]+second[i:i+1]+first[i:]
            }
        }
        if change==0||(first[n-1]==second[n-1]){
            return true
        }
        return false
    }else{
        return false
    }
}