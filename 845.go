func longestMountain(A []int) int {
	max,temp,flag:=0,1,false
	i:=1
	for ;i<len(A);i++{
		if A[i]>A[i-1]{
			break
		}
	}
	for ;i<len(A);i++{
		if flag==false{
			if A[i]>A[i-1]{
				temp++
			}else if(A[i]==A[i-1]){
				temp=1
				for ;i<len(A);i++{
					if A[i]>A[i-1]{
						break
					}
				}
			}else{
				flag=true
				temp++
				if temp>max{
					max=temp
				}
			}
		}else{
			if A[i]<A[i-1]{
				temp++
				if temp>max{
					max=temp
				}
			}else if(A[i]==A[i-1]){
				temp=2
				flag=false
				for ;i<len(A);i++{
					if A[i]>A[i-1]{
						break
					}
				}
			}else{
				flag=false
				if temp>max {
					max = temp
				}
				temp=2
			}
		}
	}
	return max
}