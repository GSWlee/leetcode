func isfinsh(i int)bool{
	switch {
	case i==2||i==1:
		return true
	case i==0:
		return false
	case i%2==1:
		return false
	default:
		return isfinsh(i/2)
	}
}

func solveSudoku(board [][]byte)  {
    var a [][]int
	for i:=0;i<9;i++{
		var tem []int
		a=append(a,tem)
		for j:=0;j<9;j++{
			a[i]=append(a[i],511)
		}
	}

	var update func(a [][]int,val byte,i,j int) ([]int,[][]int,bool)
	update=func(a [][]int,val byte,i,j int)([]int,[][]int,bool){
		nine,row,col:=511,511,511
		var taget []int
		temp := int(val)-49
		value:=int(math.Pow(float64(2),float64(temp)))
		taget=append(taget,value)
		a[i][j]=value
		for m:=0;m<9;m++{
			if a[m][j]&value!=0&&m!=i{
				a[m][j]=a[m][j]-value
				taget=append(taget,m*10+j)
			}
			if isfinsh(a[m][j]){
				if a[m][j]&col!=0{
					col=col-a[m][j]
				}else {
					return taget,a,false
				}
			}
			if a[m][j]==0{
				return  taget,a,false
			}
			if a[i][m]&value!=0&&m!=j{
				a[i][m]=a[i][m]-value
				taget=append(taget,i*10+m)
			}
			if isfinsh(a[i][m]) {
				if a[i][m]&row != 0 {
					row = row - a[i][m]
				} else {
					return taget, a, false
				}
			}
			if a[i][m]==0{
				return  taget,a,false
			}
		}
		p,q:=i/3,j/3
		for m :=p*3;m<(p*3+3);m++{
			for n := q*3;n<(q*3+3);n++{
				if a[m][n]&value!=0&&!(m==i&&n==j){
					a[m][n]=a[m][n]-value
					taget=append(taget,m*10+n)
				}
				if isfinsh(a[m][n]){
					if a[m][n]&nine!=0{
						nine=nine-a[m][n]
					}else {
						return taget,a,false
					}
				}
				if a[m][n]==0{
					return  taget,a,false
				}
			}
		}
		return taget,a,true
	}

	var reset func(a [][]int,b []int)[][]int
	reset=func(a [][]int,b []int)[][]int{
		nums:=len(b)
		value:=b[0]
		for i:=1;i<nums;i++{
			a[b[i]/10][b[i]%10]=a[b[i]/10][b[i]%10]+value
		}
		return a
	}

	var process func(a [][]int,start int)(bool,[][]int)
	process=func(a [][]int,start int)(bool,[][]int){
		i,j:=start/10,start%10
		q:=false
		for ;i<9;i++{
			for ;j<9;j++{
				if !isfinsh(a[i][j]){
					q=true
					break
				}
			}
			if q{
				break
			}
			j=0
		}
		if !q{
			return true,a
		}

		temp:=a[i][j]
		var change []int
		var flag bool
		cons:=temp
		for q:=256;q>0;q=q/2{
			if q&a[i][j]!=0 {
				temp = temp - q
				a[i][j]=q
				t:=math.Log2(float64(q))
				change,a,flag=update(a,byte(t+49),i,j)
				if flag==false{
					a[i][j]=temp
					a=reset(a,change)
					continue
				}
				flag,a=process(a,i*10+j)
				if flag{
					return flag,a
				}else{
					a[i][j]=temp
					a=reset(a,change)
				}
				if temp == 0 {
					a[i][j] = cons
					return false, a
				}
			}
		}
		a[i][j] = cons
		return false,a
	}

    for i:=0;i<9;i++{
		for j:=0;j<9;j++{
            if board[i][j]!='.'{
                _,a,_=update(a,board[i][j],i,j)
            }
			
		}
	}

    _,a=process(a,0)

    for i:=0;i<9;i++{
		for j:=0;j<9;j++{
            t:=math.Log2(float64(a[i][j]))
			board[i][j]=byte(49+t)
		}
	}

}