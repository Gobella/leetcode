package basic

import (
	"strings"
)

func RomanToInt(s string) int {
	if s==""{
		return 0
	}
	table:=make(map[string]int,0)
	table["I"]=1
	table["X"]=10
	table["V"]=5
	table["C"]=100
	table["D"]=500
	table["L"]=50
	table["M"]=1000

	b:=[]byte(s)
	lent:=len(b)
	result:=table[strings.ToUpper(string(b[lent-1]))]
	right:=result
	for i:=len(b)-1;i>0;i--{
		c:=strings.ToUpper(string(b[i-1]))
		left:=table[c]
		if right>left{
			result=result-left
		}else if right<left{
			result=result+left
		}else{
			result=result+left
		}
		right=left

	}
	return result
}