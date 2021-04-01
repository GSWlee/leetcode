func clumsy(N int) int {
    var ans int
    for i:=0;N>0&&i<4;i++{
        switch i{
            case 0:
                ans=N
            case 1:
                ans=ans*N
            case 2:
                ans=ans/N
            case 3:
                ans=ans+N
        }
        N--
    }
    for N>3{
        temp:=N*(N-1)/(N-2)
        ans=ans-temp+N-3
        N=N-4
    }
    switch N{
        case 0:
            return ans
        case 1:
            return ans-1
        case 2:
            return ans-2
        case 3:
            return ans-6
    }
    return ans
}