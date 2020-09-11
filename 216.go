ndex ,k int) bool {
	count := 0
	for i:=0;i<9;i++{
			if index%2 == 1{
					count++
			}
			index=index >> 1
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