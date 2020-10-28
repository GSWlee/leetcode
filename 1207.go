func uniqueOccurrences(arr []int) bool {
    temp:=make([]int,2001)
    for _,i:=range(arr){
        temp[i+1000]++
    }
    flag:=make([]bool,1001)
    for i,_:=range(temp){
        if flag[temp[i]]{
            return false
        }else{
            if temp[i]!=0{
                flag[temp[i]]=true
            }           
        }
    }
    return true
}