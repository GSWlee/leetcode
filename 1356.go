func sortByBits(arr []int) []int {
    var count func(a int)int
    count=func(a int)int{
        if a==0{
            return 0
        }
        ans:=1
        for a!=1{
            if a%2==1{
                ans++
                a=(a-1)/2
            }else{
                a=a/2
            }
        }
        return ans
    }
    tag:=make([]int,len(arr))
    for i,_:=range(tag){
        tag[i]=count(arr[i])
    }
    
    for i:=len(arr);i>0;i--{
        for j:=0;j<i-1;j++{
            if tag[j]>tag[j+1]{
                tempt,tempa:=tag[j],arr[j]
                tag[j]=tag[j+1]
                arr[j]=arr[j+1]
                tag[j+1]=tempt
                arr[j+1]=tempa
            }else if tag[j]==tag[j+1]{
                if arr[j]>arr[j+1]{
                    tempt,tempa:=tag[j],arr[j]
                    tag[j]=tag[j+1]
                    arr[j]=arr[j+1]
                    tag[j+1]=tempt
                    arr[j+1]=tempa
                }
            }else{
                continue
            }
        }
    }
    return arr
}