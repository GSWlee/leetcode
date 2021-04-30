type ydy struct{
    score int
    index int
}

type ydys []ydy

func(y ydys)Len()int{
    return len(y)
}

func(y ydys)Swap(i,j int){
    y[i],y[j]=y[j],y[i]
}

func(y ydys)Less(i,j int)bool{
    return y[i].score>y[j].score
}
func findRelativeRanks(score []int) []string {
    var temp ydys
    for k,v:=range score{
        temp=append(temp,ydy{v,k})
    }
    sort.Sort(temp)
    ans:=make([]string,len(score))
    for k,v:=range temp{
        if k==0{
            ans[v.index]="Gold Medal"
        }else if k==1{
            ans[v.index]="Silver Medal"
        }else if k==2{
            ans[v.index]="Bronze Medal"
        }else{
            ans[v.index]=strconv.Itoa(k+1)
        }
    }
    return ans
}