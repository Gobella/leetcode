package basic

import (
	"testing"
	"fmt"
)

func TestMatch(t *testing.T){
	fmt.Println(isMatch("",".*"),true)
	//fmt.Println(isMatch("a","ab*a"),false)
}
func Test22(t *testing.T){
	fmt.Println(generateParenthesis(3))
	//fmt.Println(isMatch("a","ab*a"),false)
}
func Test24(t *testing.T){
	head:=&ListNode{}
	var v *ListNode
	v=head
	for i:=0;i<4;i++{

		v.Next=&ListNode{}
		v.Val=i
		fmt.Println("addr",v)
		v=v.Next
	}
	v=head
	for i:=0;i<5;i++{
		fmt.Println("v.Val",v.Val)
		v=v.Next
	}
	fmt.Println(swapPairs(head))

	//fmt.Println(isMatch("a","ab*a"),false)
}

