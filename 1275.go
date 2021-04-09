func tictactoe(moves [][]int) string {
    var qipan [3][3]string
    for i,v :=range moves{
        if i%2==0{
            qipan[v[0]][v[1]]="A"
        }else{
            qipan[v[0]][v[1]]="B"
        }
    }
    for i:=0;i<3;i++{
        if qipan[i][0]==qipan[i][1]&&qipan[i][2]==qipan[i][1]&&qipan[i][1]!=""{
            return qipan[i][0]
        }
        if qipan[0][i]==qipan[1][i]&&qipan[2][i]==qipan[1][i]&&qipan[1][i]!=""{
            return qipan[0][i]
        }
    }
    if qipan[0][0]==qipan[1][1]&&qipan[1][1]==qipan[2][2]&&qipan[0][0]!=""{
        return qipan[0][0]
    }
    if qipan[2][0]==qipan[1][1]&&qipan[1][1]==qipan[0][2]&&qipan[1][1]!=""{
        return qipan[1][1]
    }
    if len(moves)==9{
        return "Draw"
    }
    return "Pending"
}