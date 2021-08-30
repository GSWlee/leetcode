type Solution struct {
    total int
    arr []int
}


func Constructor(w []int) Solution {
    sum:=0
    for _ ,v := range w{
        sum+=v
    }
    ans:=Solution{
        total:sum,
        arr:w}
    return ans
}


func (this *Solution) PickIndex() int {
    target:=rand.Intn(this.total)
    now:=0
    for k,v:=range this.arr{
        if now+v>target{
            return k
        }
        now+=v
    }
    return len(this.arr)-1
}
