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

//利用数组的有序性，双指针
func sortedSquares(A []int) []int {
    n:=len(A)
    i,j,pos:=0,n-1,n-1
    ans:=make([]int,n)
    for i<=j{
        if A[i]*A[i]>A[j]*A[j]{
            ans[pos]=A[i]*A[i]
            pos--
            i++
        }else{
            ans[pos]=A[j]*A[j]
            pos--
            j--
        }
    }
    return ans
}

