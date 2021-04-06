func findMinFibonacciNumbers(k int) int {
    fib:=make([]int,0)
    fib=append(fib,1)
    fib=append(fib,1)
    for true{
        temp:=fib[len(fib)-1]+fib[len(fib)-2]
        if temp<=k{
            fib=append(fib,temp)
        }else{
            break
        }
    }
    ans:=0
    for i:=len(fib)-1;i>=0;i--{
        if k==0{
            break
        }else{
            if k>=fib[i]{
                ans++
                k=k-fib[i]
            }
        }
    }
    return ans
}