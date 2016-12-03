package basic

import (
	"container/list"
	"fmt"
	"strings"
)

func LetterCombinations(digits string) []string {
	numtostring:=make(map[int][]byte)
	numtostring[2]=[]byte("abc")
	numtostring[3]=[]byte("def")
	numtostring[4]=[]byte("ghi")
	numtostring[5]=[]byte("jkl")
	numtostring[6]=[]byte("mno")
	numtostring[7]=[]byte("pqrs")
	numtostring[8]=[]byte("tuv")
	numtostring[9]=[]byte("wxyz")

	result:=make([]string,0)
	resList:=list.New()
	digits=strings.Replace(digits,"0","",-1)
	digits=strings.Replace(digits,"1","",-1)
	if digits==""{
		return result
	}

	item:=make([]byte,len(digits))
	digitsb:=[]byte(digits)
	tree(0,item,digitsb,resList,numtostring)
	for e := resList.Front(); e != nil; e = e.Next(){
		str,_:=e.Value.(string)
		result=append(result,str)
	}
	return result
}

func tree(itemIndex int,item []byte,digits []byte,result *list.List,mapp map[int][]byte){
	start:=byte('0')
	if itemIndex==len(digits){
		fmt.Println(string(item))
		result.PushBack(string(item))
		return
	}
	b:=digits[itemIndex]
	c:=int(b-start)

	for _,v :=range mapp[c]{
		//fmt.Println("c",v)
		item[itemIndex]=v
		tree(itemIndex+1,item,digits,result,mapp)
	}

}