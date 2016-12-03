package basic

import (
	"testing"
	"fmt"
)

func TestInsertSort(t *testing.T) {
	a:=[]int{4,5,3,2,1}
	c:=SelectSort(a,5)
	fmt.Println("c:",c[0],c[1],c[2],c[3],c[4])
}

func TestConvert(t *testing.T) {
	fmt.Println(Convert("AB",1))
}