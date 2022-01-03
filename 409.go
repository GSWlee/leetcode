func longestPalindrome(s string) int {
    dict:=map[rune]int{}
    sum:=0
    for _,v:=range s{
        dict[v]+=1
        if dict[v]%2==1{
            sum++
        }else{
            sum--
        }
    }
    if sum==0{
        return len(s)
    }
    return len(s)-sum+1
}