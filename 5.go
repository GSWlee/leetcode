func longestPalindrome(s string) string {
    start,end:=0,0
    for i:=0;i<len(s);i++{
        s1,e1:=findedge(s,i,i)
        s2,e2:=findedge(s,i,i+1)
        if e1-s1>end-start{
            start,end=s1,e1
        }
        if e2-s2>end-start{
            start,end=s2,e2
        }
    }
    return s[start:end+1]
}

func findedge(s string,i int ,j int)(int,int){
    for;i>=0&&j<len(s)&&s[i]==s[j];i,j=i-1,j+1{}
    return i+1,j-1
}