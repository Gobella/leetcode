package basic

import "fmt"

func Convert(s string, numRows int) string {

	str:=[]byte(s)
	len0 := len(str)

	if len0<=numRows||numRows<=1{
		return s
	}

	step:=2*numRows-2
	newbyte:=[]byte(s)
	c:=0

	for row:=0;row<numRows;row++{
		for index:=row;index<len0;index=index+step{
			fmt.Println(c,index)
			newbyte[c]=str[index]
			c++
			if row==0||row==numRows-1{
				continue
			}

			interval:=index+2*(numRows-row-1)
			if interval<len0{
				fmt.Println(c,interval)
				newbyte[c]=str[interval]
				c++
			}
		}

	}
	return string(newbyte)

}