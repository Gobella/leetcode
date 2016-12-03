package basic

import (
	"container/list"
)

func IsValid(s string) bool {
	stack:=list.New()
	mapp:=make(map[string]string)
	mapp["{"]="}"
	mapp["["]="]"
	mapp["("]=")"

	mappc:=make(map[string]string)
	mappc["}"]="{"
	mappc["]"]="["
	mappc[")"]="("

	b:=[]byte(s)
	for _,v:=range b{
		char:=string(v)
		if _,ok:=mapp[char];ok{
			stack.PushFront(char)
			continue
		}
		if value,ok:=mappc[char];ok{
			e:=stack.Front()
			if e==nil{
				return false
			}
			c,okk:=e.Value.(string)
			if !okk||value!=c{
				return false
			}
			stack.Remove(stack.Front())
		}
	}
	if stack.Len()!=0{
		return false
	}
	return true

}
