func numDecodings(s string) int {
    dp:=make([]int,len(s)+1)
    for i:=len(s)-1;i>=0;i--{
        if s[i:i+1]=="0"{
            dp[i]=0
        }else{
            if dp[i+1]!=0{
                dp[i]=dp[i+1]
                if ans,err:=strconv.Atoi(s[i:i+2]);err==nil&&ans<=26{
                    if dp[i+2]==0&&i+2==len(s){
                        dp[i]++
                    }else{
                        dp[i]=dp[i+2]+dp[i]
                    }
                }
            }else{
                if i+1==len(s){
                    dp[i]=1
                }else{
                    if ans,err:=strconv.Atoi(s[i:i+2]);err==nil&&ans<=26{
                        if dp[i+2]==0&&i+2==len(s){
                            dp[i]=1
                        }else if dp[i+2]==0&&i+2!=len(s){
                            break
                        }else{
                            dp[i]=dp[i+2]
                        }
                    }else{
                        break
                    }
                }
            }
        }
    }
    return dp[0]
}