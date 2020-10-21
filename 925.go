func isLongPressedName(name string, typed string) bool {
    i,j:=0,0
    pre:=name[0]
    for i<len(name)&&j<len(typed){
        if name[i]==typed[j]{
            pre=name[i]
            i++
            j++
        }else if typed[j]==pre{
            j++
        }else{
            return false
        }
    }
    if i!=len(name){
        return false
    }else{
        for j<len(typed){
            if typed[j]!=pre{
                return false
            }
            j++
        }
        return true
    }
}