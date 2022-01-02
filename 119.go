// 递归
func getRow(rowIndex int) []int {
    if rowIndex==0{
        return []int{1}
    }else{
        temp:=getRow(rowIndex-1)
        ans:=[]int{1}
        for k,_:=range temp{
            if k>0{
                ans=append(ans,temp[k]+temp[k-1])
            }
        }
        ans=append(ans,1)
        return ans
    }
}

//减少内存开销
func getRow(rowIndex int) []int {
    ans:=[]int{1}
    for i:=1;i<=rowIndex;i++{
        temp:=[]int{1}
        for k,_:=range ans{
            if k>0{
                temp=append(temp,ans[k]+ans[k-1])
            }
        }
        temp=append(temp,1)
        ans=temp
    }
    return ans
}