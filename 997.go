//直接插入排序
func sortedSquares(A []int) []int {
    for i:=0;i<len(A);i++{
        A[i]=A[i]*A[i]
        for j:=i;j>0;j--{
            if A[j]<A[j-1]{
                temp:=A[j]
                A[j]=A[j-1]
                A[j-1]=temp
            }else{
                break
            }
        }        
    }
    return A
}