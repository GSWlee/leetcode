func canFormArray(arr []int, pieces [][]int) bool {
    flag:=false
    var isin func(arr,target []int)bool
    isin=func(arr,target []int)bool{
        if len(target)>len(arr){
            return false
        }else{
            for i,_:=range(target){
                if arr[i]!=target[i]{
                    return false
                }
            }
        }
        return true
    }
    for _,j :=range(pieces){
        if isin(arr,j){
            if len(arr)==len(j){
                return true
            }
            flag=canFormArray(arr[len(j):], pieces)
            if flag==true{
                return true
            }
        }
    }
    return false
}