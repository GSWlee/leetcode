func judgeSquareSum(c int) bool {
    num:=int(math.Sqrt(float64(c/2)))
    for i:=0;i<=num+1;i++{
        if ans:=math.Sqrt(float64(c-i*i));ans==float64(int(ans)){
            return true
        }
    }
    return false
}