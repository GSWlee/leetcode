func max(a int,b int) int{
    if a<b{
        return b
    }
    return a
}

func maxProfit(k int,prices []int) int {
    length:=len(prices)
    if length==0||k==0{
        return 0
    }
    last:=make([]int,k*2+1)
    last[1]=-prices[0]
    for i:=2;i<2*k+1;i++{
        last[i]=int(math.NaN())
    }
    for i:=1;i<length;i++{
        temp:=make([]int,k*2+1)
        for j:=1;j<=k;j++{
            if last[2*j-2]!=int(math.NaN()){
                if last[2*j-1]!=int(math.NaN()){
                    temp[2*j-1]=max(last[2*j-2]-prices[i],last[2*j-1])
                }else{
                    temp[2*j-1]=last[2*j-2]-prices[i]
                }
            }else{
                temp[2*j-1]=int(math.NaN())
            }
            if last[2*j-1]!=int(math.NaN()){
                if last[2*j]!=int(math.NaN()){
                    temp[2*j]=max(last[2*j-1]+prices[i],last[2*j])
                }else{
                    temp[2*j]=last[2*j-1]+prices[i]
                }
            }else{
                temp[2*j]=int(math.NaN())
            }
        }
        last=temp
    }
    MAX:=0
    for i:=1;i<2*k+1;i++{
        if last[i]>MAX{
            MAX=last[i]
        }
    }
    return MAX
}