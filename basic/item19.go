package basic


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count:=0
	mapp:=make(map[int]*ListNode)
	thead:=head

	for ;head!=nil;{
		count++
		mapp[count]=head
		head=head.Next
	}
	if count==n{
		thead=thead.Next
	}else if n==1{
		mapp[count-n].Next=nil
	}else{
		mapp[count-n].Next=mapp[count-n+2]
	}

	return thead
}

//double pointer

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	count:=1
	dis:=0
	point1:=head
	point2:=head
	for ;point2.Next!=nil;{
		count++
		point2=point2.Next
		if dis==n{
			point1=point1.Next
		}else{
			dis++
		}

	}
	if count==n{
		head=point1.Next
	}else{
		point1.Next=point1.Next.Next
	}

	return head
}