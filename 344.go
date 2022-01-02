func reverseString(s []byte)  {
    st,e:=0,len(s)-1
    for st<e{
        temp:=s[st]
        s[st]=s[e]
        s[e]=temp
        st++
        e--
    }
}