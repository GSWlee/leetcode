func canCompleteCircuit(gas []int, cost []int) int {
    for i:=0;i<len(gas);i++{
        sum:=0
        flag:=true
        for j:=0;j<len(gas);j++{
            temp:=(i+j)%len(gas)
            sum=sum+gas[temp]
            sum=sum-cost[temp]
            if sum<0{
                flag=false
                break
            }
        }
        if flag==true{
            return i
        }
    }
    return -1
}