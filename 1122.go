func relativeSortArray(arr1 []int, arr2 []int) []int {
    tag:=make([]int,len(arr2)+1)
    for _,i:=range(arr1){
        j:=0
        for ;j<len(arr2);j++{
            if arr2[j]==i{
                break
            }
        }
        tag[j]++
    }
    sum:=tag[len(tag)-1]
    for i:=1;i<len(tag);i++{
        tag[i]=tag[i]+tag[i-1]
    }
    ans:=make([]int,len(arr1))
    for i:=len(arr1)-1;i>=0;i--{
        j:=0
        for ;j<len(arr2);j++{
            if arr2[j]==arr1[i]{
                break
            }
        }
        ans[tag[j]-1]=arr1[i]
        tag[j]--
    }
    for ;sum>0;sum--{
        for j:=len(ans)-1;j>len(ans)-sum;j--{
            if ans[j]<ans[j-1]{
                temp:=ans[j]
                ans[j]=ans[j-1]
                ans[j-1]=temp
            }
        }
    }
    return ans
}