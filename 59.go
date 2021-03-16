func generateMatrix(n int) [][]int {
	ans:=make([][]int,n)
	for i:=range(ans){
		ans[i]=make([]int,n)
	}
	flag:=0
	q:=1
	for q<=n*n{
		switch flag%4{
		case 0:
			start:=flag/4
			for i:=start;i<n-start;i++{
				ans[start][i]=q
				q++
			}
			flag++
		case 1:
			start:=flag/4
			for i:=start+1;i<n-start-1;i++{
				ans[i][n-1-start]=q
				q++
			}
			flag++
		case 2:
			start:=flag/4
			for i:=n-1-start;i>=start;i--{
				ans[n-start-1][i]=q
				q++
			}
			flag++
		case 3:
			start:=flag/4
			for i:=n-2-start;i>start;i--{
				ans[i][start]=q
				q++
			}
			flag++
		}
	}
	return ans
}