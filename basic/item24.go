package basic


func swapPairs(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	thead:=head.Next

	for ;head!=nil&&head.Next!=nil;{
		temp:=head.Next.Next
		head.Next.Next=head
		if temp!=nil&&temp.Next!=nil{
			head.Next=temp.Next
		}else{
			head.Next=temp
		}

		head=temp
	}
	return  thead
}