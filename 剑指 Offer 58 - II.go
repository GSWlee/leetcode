func reverseLeftWords(s string, n int) string {
    var reverse func(s string,start,end int)string
    reverse=func(s string,start,end int)string{
        b := []rune(s)
        for start<end{
            temp:=b[start]
            b[start]=b[end]
            b[end]=temp
            start++
            end--
        }
        s=string(b)
        return s 
    }
    s=reverse(s,0,len(s)-1)
    s=reverse(s,0,len(s)-n-1)
    s=reverse(s,len(s)-n,len(s)-1)
    return s
}