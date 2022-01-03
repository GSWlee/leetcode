func checkInclusion(s1 string, s2 string) bool {
    n,m:=len(s1),len(s2)
    if n>m{
        return false
    }
    cnt:=[26]int{}
    for _,ch:=range s1{
        cnt[ch-'a']--
    }
    left:=0
    for right,v:=range s2{
        x:=v-'a'
        cnt[x]++
        for cnt[x]>0{
            cnt[s2[left]-'a']--
            left++
        }
        if right-left+1==n{
            return true
        }
    }
    return false
}