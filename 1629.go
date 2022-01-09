func slowestKey(releaseTimes []int, keysPressed string) byte {
    s:=0
    ans:=keysPressed[0]
    times:=0
    for k,v:=range releaseTimes{
        temp:=keysPressed[k]
        t:=v-s
        if t>times{
            ans=temp
            times=t
        }else if t==times{
            if temp>ans{
                ans=temp
            }
        }
        s=v
    }
    return ans
}