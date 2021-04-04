func numRabbits(answers []int) int {
    m := make(map[int]int)
    ans:=0
    for _,v:=range answers{
        if v==0{
            ans++
            continue
        }
        if _,ok:=m[v];ok{
            m[v]++
        }else{
            m[v]=1
        }
    }
    
    for k,v:=range m{
        for v>k+1{
            v=v-k-1
            ans=ans+k+1
        }
        ans=ans+k+1
    }
    return ans
}