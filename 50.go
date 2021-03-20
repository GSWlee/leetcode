func myPow(x float64, n int) float64 {
	if n==0{
		return 1.00000
	}
	flag:=true
	if n<0{
		n=-n
		flag=false
	}
	i:=0
	for j:=1;j<=n;j=j*2{
		i++
	}
	dp:=make([]float64,i)
	dp[0]=x
	for j:=1;j<i;j++{
		dp[j]=dp[j-1]*dp[j-1]
	}
	ans:=1.00000
	for n!=0{
		temp:=int(math.Pow(2.00000,float64(i)))
		if n%temp==0{
			ans=ans*dp[i]
			break
		}else{
			if n>temp{
				ans*=dp[i]
				n=n-temp
			}
		}
		i--
	}
	if flag{
		return ans
	}
	return 1/ans
}