func strStr(haystack string, needle string) int {
    len1,len2:=len(haystack),len(needle)
    if len2==0{
        return 0
    }
    value:=make([]int,len2)
    for i:=1;i<len2;i++{
        if needle[i]==needle[value[i-1]]{
            value[i]=value[i-1]+1
        }else{
            for j:=value[i-1];j>0;j--{
                m:=0
                for ;m<j;m++{
                    if needle[i-m]!=needle[j-m-1]{
                        break
                    }
                }
                if m==j{
                    value[i]=j
                    break
                }
            }
        }
    }
    for j:=len2-1;j>0;j--{
        value[j]=value[j-1]
    } 
    value[0]=-1
    now:=0
    i:=0
    for ;i<len1&&now<len2;i++{
        if haystack[i]==needle[now]{
            now++
        }else{
            if now!=0{
                now=value[now]
                i--
            }
        }
    }
    if now==len2{
        return i-now
    }
    return -1
}