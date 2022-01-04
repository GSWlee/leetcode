//双hash
func wordPattern(pattern string, s string) bool {
    dict:=map[string]byte{}
    dict2:=map[byte]int{}
    index,left:=0,0
    for k,v:= range s{
        if v==' '{
            temp:=s[left:k]
            if index==len(pattern){
                return false
            }
            if _,ok:=dict[temp];ok{
                if dict[temp]==pattern[index]{
                    index++
                    left=k+1
                    continue
                }else{
                    return false
                }
            }else{
                if _,ok:=dict2[pattern[index]];ok{
                    return false
                }else{
                    dict2[pattern[index]]=1
                    dict[temp]=pattern[index]
                    index++
                    left=k+1
                }
            }
        }
    }
    if left<len(s){
        temp:=s[left:]
        if index==len(pattern){
            return false
        }
        if _,ok:=dict[temp];ok{
                if dict[temp]==pattern[index]{
                    index++
                }else{
                    return false
                }
            }else{
                if _,ok:=dict2[pattern[index]];ok{
                    return false
                }else{
                    dict2[pattern[index]]=1
                    dict[temp]=pattern[index]
                    index++
                }
            }
    }
    if index<len(pattern){
        return false
    }
    return true
}