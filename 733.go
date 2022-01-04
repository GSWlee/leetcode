func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	width,high:=len(image[0]),len(image)
	flag:=[][]int{}
	for i:=0;i<high;i++{
		flag=append(flag,make([]int,width))
	}
	stack:=[][2]int{{sr,sc}}
    tag:=image[sr][sc]
	for len(stack)>0{
		x,y:=stack[0][0],stack[0][1]
		if image[x][y]!=tag{
			stack=stack[1:]
		}else{
			image[x][y]=newColor
			flag[x][y]=1
			if x-1>=0 &&image[x-1][y]==tag&&flag[x-1][y]==0{
				stack=append(stack,[2]int{x-1,y})
			}
			if x+1<high &&image[x+1][y]==tag&&flag[x+1][y]==0{
				stack=append(stack,[2]int{x+1,y})
			}
			if y+1<width &&image[x][y+1]==tag&&flag[x][y+1]==0{
				stack=append(stack,[2]int{x,y+1})
			}
			if y-1>=0 &&image[x][y-1]==tag&&flag[x][y-1]==0{
				stack=append(stack,[2]int{x,y-1})
			}
			stack=stack[1:]
		}
	}
	return image
}