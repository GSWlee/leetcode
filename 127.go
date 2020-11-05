func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList=append(wordList,beginWord)
	shortpath,flag:=make([]int,len(wordList)),make([]bool,len(wordList))
	var diff func(str1,str2 string)int
	diff=func(str1,str2 string)int{
		sum:=0
		for i,_:=range(str1){
			if str1[i]!=str2[i]{
				sum++
			}
		}
		return sum
	}
	var finsh func(flag []bool)bool
	finsh=func(flag []bool)bool{
		for _,i:=range(flag){
			if i==false{
				return true
			}
		}
		return false
	}
	var arr [][]int
	for i:=0;i<len(wordList);i++{
		temp:=make([]int,len(wordList))
		for j:=0;j<len(wordList);j++{
			if diff(wordList[i],wordList[j])<=1{
				temp[j]=diff(wordList[i],wordList[j])
			}else{
				temp[j]=10000
			}
		}
		arr=append(arr,temp)
	}

	for i,_:=range(shortpath){
		if i==len(shortpath)-1{
			shortpath[i]=1
		}else{
			shortpath[i]=10000
		}
	}
	for finsh(flag){
		min:=0
		for i:=0;i<len(flag);i++{
			if flag[i]==false{
				min=i
				break
			}
		}
		for i:=0;i<len(shortpath);i++{
			if flag[i]==false&&shortpath[i]<shortpath[min]{
				min=i
			}
		}
		flag[min]=true
		for i:=0;i<len(shortpath);i++{
			if arr[min][i]+shortpath[min]<shortpath[i]{
				shortpath[i]=arr[min][i]+shortpath[min]
			}
		}
	}
	for i:=0;i<len(wordList)-1;i++{
		if diff(endWord,wordList[i])==0{
            if shortpath[i]==10000{
                return 0
            }
			return shortpath[i]
		}
	}
	return 0
}