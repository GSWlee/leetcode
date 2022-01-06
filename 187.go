func findRepeatedDnaSequences(s string) (ans []string) {
    dict:=map[string]int{}
    if len(s)<=10{
        return 
    }
    dict[s[:10]]=1
    for i:=1;i+10<=len(s);i++{
        if v,ok:=dict[s[i:i+10]];ok{
            if v==1{
                ans=append(ans,s[i:i+10])
                dict[s[i:i+10]]++
            }   
        }else{
            dict[s[i:i+10]]=1
        }
    }
    return
}