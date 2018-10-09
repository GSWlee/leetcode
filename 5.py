"""
最长回文子串
manachar算法
"""
def findall(s,i):
    if i=0 or i=len(s)-1:
        return 1
    else:
        q=0
        while(q<i and q< len(s)-i):
            if s[i-q]==s[i+q]:
                q=q+1
            else:
                return q

def findside(s,i,)
s="ababa"
q=list(s)
s1='#'+"#".join(q)+'#'
fz=list()
pos=0
max=0
for i in range(0,len(s1)):
    if i >= max:
        q=findall(s1,i)
        fz.append(q)
    else:
        if fz[2*pos-i]+i<=max:
            fz.append(fz[2*pos-i])
        else:
            q=findside(s1,i,)
            fz.append(q)
max=0
for i in fz:
    max=max if fz[max]>fz[i] else i

ans=s1[i-fz[i]:i+fz[i]]