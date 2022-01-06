func simplifyPath(path string) string {
    dir:=[]string{}
    start:=-1
    for i:=0;i<len(path);i++{
        if path[i]=='/'&&start!=-1{
            if path[start:i]==".."{
                if len(dir)>0{
                    dir=dir[:len(dir)-1]
                }else{
                    start=-1
                }
            }else if path[start:i]=="."{
                start=-1
            }else{
                dir=append(dir,path[start:i])
            }
            start=-1
        }else if path[i]=='/'&&start==-1{
            continue
        }else if path[i]!='/'&&start==-1{
            start=i
        }else{
            continue
        }
    }
    if start!=-1{
        i:=len(path)
        if path[start:i]==".."{
            if len(dir)>0{
                dir=dir[:len(dir)-1]
            }else{
                start=-1
            }
        }else if path[start:i]=="."{
            start=-1
        }else{
            dir=append(dir,path[start:i])
        }
    }
   if len(dir)==0{
       return "/"
   }else{
       var ans string
       for _,v:=range dir{
           ans=ans+string("/")+v
       }
       return ans
   }
}