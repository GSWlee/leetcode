func findRotateSteps(ring string, key string) int {
	dp:=make([]int,len(ring))
	var distant func(i,point int)int
	distant=func(i,point int)int{
		s,n:=i-point,point-i
		if s<0{
			s=s+len(ring)
		}
		if n<0{
			n=n+len(ring)
		}
		if s<n{
			return s
		}else{
			return n
		}
	}
	var points=[]int{0}
	var new=[]int{}
	for i,_:=range(key){
		for j:=0;j<len(ring);j++{
			if ring[j]==key[i]{
				new=append(new,j)
				min:=10000
				for _,m:=range(points){
					temp:=distant(j,m)+1+dp[m]
					if min>=temp{
						min=temp
					}
				}
                dp[j]=min
			}
		}
		points=new
		new=[]int{}
	}
	min:=10000
	for j:=0;j<len(ring);j++{
		if ring[j]==key[len(key)-1]{
			if dp[j]<min{
				min=dp[j]
			}
		}
	}
	return min
}