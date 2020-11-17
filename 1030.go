func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
    var tag []int
    var ans [][]int
    var dis func(a,b,c,d int)int
    dis=func(a,b,c,d int)int{
        sum:=0
        if a-c>0{
            sum=sum+a-c
        }else{
            sum=sum+c-a
        }
        if b-d>0{
            sum=sum+b-d
        }else{
            sum=sum+d-b
        }
        return sum
    }
    for i:=0;i<R;i++{
        for j:=0;j<C;j++{
            tag=append(tag,dis(i,j,r0,c0))
            ans=append(ans,[]int{i,j})
        }
    }
    for i:=len(ans)-1;i>0;i--{
        for j:=0;j<i;j++{
            if tag[j]>tag[j+1]{
                tempd,tempp:=tag[j],ans[j]
                tag[j],ans[j]=tag[j+1],ans[j+1]
                tag[j+1],ans[j+1]=tempd,tempp   
            }
        }
    }
    return ans
}