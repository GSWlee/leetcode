func shipWithinDays(weights []int, D int) int {
	var comp func(j int) int
	comp=func(j int) int{
		day,temp:=0,0
		for _,v:=range weights{
			if temp+v<=j{
				temp=temp+v
			}else{
				day++
				if v>j{
					return -1
				}else{
					temp=v
				}
			}
		}
		if temp!=0{
			day++
		}
		return day
	}
	sum:=0
	for _,v:=range weights{
		sum+=v
	}
	var find func(start int,end int) int
	find=func(start int,end int) int{
		if start>end{
			return -1
		}else{
			now:=(start+end)/2
			dayN:=comp(now)
			dayM:=comp(now-1)
			if dayN!=-1&&dayN<=D&&(dayM>D||dayM==-1){
				return now
			}else if dayN>D ||dayN==-1{
				return find(now+1,end)
			}else{
				return find(start,now-1)
			}
		}
	}
	return find(1,sum)
}