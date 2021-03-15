func spiralOrder(matrix [][]int) []int {
	ans:=[]int{}
	flag:=0
	sum:=len(matrix)*len(matrix[0])
	m,n:=len(matrix),len(matrix[0])
	for len(ans)!=sum{
		switch flag%4{
		case 0:
			start:=flag/4
			for i:=start;i<n-start;i++{
				ans=append(ans,matrix[start][i])
			}
			flag++
		case 1:
			start:=flag/4
			for i:=start+1;i<m-start-1;i++{
				ans=append(ans,matrix[i][n-1-start])
			}
			flag++
		case 2:
			start:=flag/4
			for i:=n-1-start;i>=start;i--{
				ans=append(ans,matrix[m-start-1][i])
			}
			flag++
		case 3:
			start:=flag/4
			for i:=m-2-start;i>start;i--{
				ans=append(ans,matrix[i][start])
			}
			flag++
		}
	}
	return ans
}