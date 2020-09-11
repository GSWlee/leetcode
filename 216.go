func combinationSum3(k int, n int) [][]int {
	var arr [][]int
	nums :=[9]int  {1,2,3,4,5,6,7,8,9}
	var index=math.Pow(2,float64(k))-1
	end := math.Pow(2,9)-math.Pow(2,float64(9-k))
	begin :=int(index)
	end1 :=int(end)
	for begin<=end1{
		if isK(begin ,k){
			var temp []int
			sum :=0
			for i:=0;i<9;i++{
				if (begin & (1<<i))==(1<<i) {
					temp = append(temp, nums[i])
					sum+=nums[i]
				}
			}
			if sum == n{
				arr = append(arr, temp)
			}

		}
		begin++
	}
	return  arr
}

func isK(index ,k int) bool {
	count := 0
	for i:=0;i<9;i++{
		if index%2 == 1{
			count++
		}
		index=index >> 1
	}
	return count==k
}