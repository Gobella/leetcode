package basic

import (
	"container/list"
	"fmt"
)

func generateParenthesis(n int) []string {
	result:=make([]string,0)
	resultList:=list.New()
	printOne(n,n,"",resultList)

	for e := resultList.Front(); e != nil; e = e.Next() {
		v,_:=e.Value.(string)
		result=append(result,v)
	}
	return result
}

func printOne(left int,right int,s string,result *list.List){
	fmt.Println(s)
	if left==0&&right==0{

		result.PushBack(s)
	}
	if left>0{
		printOne(left-1,right,s+"(",result)
	}
	if right>0&&left<right{
		printOne(left,right-1,s+")",result)
	}
}
