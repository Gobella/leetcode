package basic

import (
	"fmt"
	"strconv"
)

func Reverse(x int) int {
	var c int32=1
	c=c << 31
	min:=c
	max:=^c
	negative:=false
	if x<0{
		negative=true
		x=-x
	}
	str := fmt.Sprintf("%d",x)
	by:=[]byte(str)
	newInt:=[]byte(str)
	for i:=len(by)-1;i>-1;i--{
		newInt[len(by)-i-1]=by[i]
	}

	b,error := strconv.Atoi(string(newInt))
	if error!=nil{
		return 0
	}
	if negative{
		b=-b
	}
	if b>int(max)||b<int(min){
		return 0
	}
	return b
}
