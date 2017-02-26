package leetcode;

public class Item203 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		ListNode l=new ListNode(1);
		ListNode head=l;
//		l.next=new ListNode(1);
//		l=l.next;
//		l.next=new ListNode(3);
//		l=l.next;
//		l.next=new ListNode(3);
		ListNode ls=removeElements(head,1);
		while(ls!=null){
			System.out.println(ls.val);
			ls=ls.next;
		}
	}
	public static ListNode removeElements(ListNode head, int val) {
		if(head==null){
			return null;
		}
		while(head!=null&&head.val==val){
			head=head.next;
		}
		if(head ==null)
			return head;
		ListNode cusor=head.next;
		ListNode beforeCusor=head;
		while(cusor!=null){
			if(cusor.val==val){
				beforeCusor.next=cusor.next;
			}else{
				beforeCusor=beforeCusor.next;				
			}	
			cusor=cusor.next;
		}
        return head;
    }

}
