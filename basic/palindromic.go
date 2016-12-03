package basic


func LongestPalindrome(s string) string {
	b:=[]byte(s)
	if len(b)<=1{
		return s
	}

	if len(b)==2&&b[0]!=b[1]{
		return ""
	}

	var length int

	temp:=""
	tempLen:=0

	for a:=0;a<len(b)-1;{
		length=0
		p:=0
		c:=a+1
		for ;p<a+1&&c+p<len(b);p++{
			if b[a-p]==b[c+p]{
				length=length+2
			}else{
				break
			}
		}
		if tempLen<length{
			tempLen=length
			temp=string(b[a-p+1:c+p])
		}
		a++
	}
	for i:=1;i<len(b)-1;i++{
		j:=1
		length=1
		for ;j<i+1&&j+i<len(b);j++{
			if b[i-j]==b[j+i]{
				length=length+2
			}else{
				break
			}
		}
		if tempLen<length{
			tempLen=length
			temp=string(b[i-j+1:i+j])
		}
	}
	return temp
}

func LongestPalindrome1(s string) string {
	b:=[]byte(s)
	tempB:=make([]byte,0)
	for _,v:=range b{
		tempB=append(tempB,0)
		tempB=append(tempB,v)
	}
	tempB=append(tempB,0)
	Blen:=len(tempB)
	RLen:=make([]int,0)
	var maxright int=0
	var pos int=0

	for j,_:=range tempB{
		m:=0
		RLen=append(RLen,m)
		if j<maxright{
			if maxright-j>RLen[2*pos-j]{
				RLen[j]=RLen[2*pos-j]
			}else{
				RLen[j]=maxright-j
			}
		}
		for ;j+RLen[j]<Blen&&j-RLen[j]>-1&&tempB[j+RLen[j]]==tempB[j-RLen[j]];RLen[j]++{
		}

		if maxright<RLen[j]+j-1{
			maxright=RLen[j]+j-1
			pos=j
		}
		if maxright<j{
			maxright=j
			pos=j
		}
	}

	var maxLen int=0
	var center int=0
	for k,v:=range RLen{
		if v>maxLen{
			maxLen=v
			center=k
		}
	}
	maxLen--

	resultTemp:=tempB[center-maxLen:center+maxLen+1]
	result:=make([]byte,0)
	for _,v:=range resultTemp{
		if v!=0{
			result=append(result,v)
		}

	}
	return string(result)
}